package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogoproto "github.com/gogo/protobuf/proto"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// ValidateGenesis performs basic validation for the following:
// - params are valid param types with valid properties
// - proto messages are valid proto messages
// - the credit type referenced in each credit class exists
// - the credit class referenced in each project exists
// - the tradable amount of each credit batch complies with the credit type precision
// - the retired amount of each credit batch complies with the credit type precision
// - the calculated total amount of each credit batch matches the total supply
// An error is returned if any of these validation checks fail.
func ValidateGenesis(data json.RawMessage, params Params) error {
	if err := params.Validate(); err != nil {
		return err
	}

	db := dbm.NewMemDB()
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{
		JSONValidator: func(m proto.Message) error {
			return validateMsg(m)
		},
	})
	if err != nil {
		return err
	}

	ormCtx := ormtable.WrapContextDefault(backend)
	ss, err := api.NewStateStore(ormdb)
	if err != nil {
		return err
	}

	jsonSource, err := ormjson.NewRawMessageSource(data)
	if err != nil {
		return err
	}

	err = ormdb.ImportJSON(ormCtx, jsonSource)
	if err != nil {
		return err
	}

	if err := ormdb.ValidateJSON(jsonSource); err != nil {
		return err
	}

	abbrevToPrecision := make(map[string]uint32) // map of credit abbreviation to precision
	ctItr, err := ss.CreditTypeTable().List(ormCtx, &api.CreditTypePrimaryKey{})
	if err != nil {
		return err
	}
	for ctItr.Next() {
		ct, err := ctItr.Value()
		if err != nil {
			return err
		}
		abbrevToPrecision[ct.Abbreviation] = ct.Precision
	}
	ctItr.Close()

	cItr, err := ss.ClassTable().List(ormCtx, api.ClassPrimaryKey{})
	if err != nil {
		return err
	}
	defer cItr.Close()

	// make sure credit type exist for class abbreviation in params
	for cItr.Next() {
		class, err := cItr.Value()
		if err != nil {
			return err
		}

		if _, ok := abbrevToPrecision[class.CreditTypeAbbrev]; !ok {
			return sdkerrors.ErrNotFound.Wrapf("credit type not exist for %s abbreviation", class.CreditTypeAbbrev)
		}
	}

	projectKeyToClassKey := make(map[uint64]uint64) // map of project key to class key
	pItr, err := ss.ProjectTable().List(ormCtx, api.ProjectPrimaryKey{})
	if err != nil {
		return err
	}
	defer pItr.Close()

	for pItr.Next() {
		project, err := pItr.Value()
		if err != nil {
			return err
		}

		if _, exists := projectKeyToClassKey[project.Key]; exists {
			continue
		}
		projectKeyToClassKey[project.Key] = project.ClassKey
	}

	batchIdToPrecision := make(map[uint64]uint32) // map of batchID to precision
	batchDenomToIdMap := make(map[string]uint64)  // map of batchDenom to batchId
	bItr, err := ss.BatchTable().List(ormCtx, api.BatchPrimaryKey{})
	if err != nil {
		return err
	}
	defer bItr.Close()

	// create index batchID => precision for faster lookup
	for bItr.Next() {
		batch, err := bItr.Value()
		if err != nil {
			return err
		}

		batchDenomToIdMap[batch.Denom] = batch.Key

		if _, exists := batchIdToPrecision[batch.Key]; exists {
			continue
		}

		class, err := ss.ClassTable().Get(ormCtx, projectKeyToClassKey[batch.ProjectKey])
		if err != nil {
			return err
		}

		if class.Key == projectKeyToClassKey[batch.ProjectKey] {
			batchIdToPrecision[batch.Key] = abbrevToPrecision[class.CreditTypeAbbrev]
		}
	}

	batchIdToCalSupply := make(map[uint64]math.Dec) // map of batchID to calculated supply
	batchIdToSupply := make(map[uint64]math.Dec)    // map of batchID to actual supply
	bsItr, err := ss.BatchSupplyTable().List(ormCtx, api.BatchSupplyPrimaryKey{})
	if err != nil {
		return err
	}
	defer bsItr.Close()

	// calculate total supply for each credit batch (tradable + retired supply)
	for bsItr.Next() {
		batchSupply, err := bsItr.Value()
		if err != nil {
			return err
		}

		tSupply := math.NewDecFromInt64(0)
		rSupply := math.NewDecFromInt64(0)
		if batchSupply.TradableAmount != "" {
			tSupply, err = math.NewNonNegativeFixedDecFromString(batchSupply.TradableAmount, batchIdToPrecision[batchSupply.BatchKey])
			if err != nil {
				return err
			}
		}
		if batchSupply.RetiredAmount != "" {
			rSupply, err = math.NewNonNegativeFixedDecFromString(batchSupply.RetiredAmount, batchIdToPrecision[batchSupply.BatchKey])
			if err != nil {
				return err
			}
		}

		total, err := math.SafeAddBalance(tSupply, rSupply)
		if err != nil {
			return err
		}

		batchIdToSupply[batchSupply.BatchKey] = total
	}

	// calculate credit batch supply from genesis tradable, retired and escrowed balances
	if err := calculateSupply(ormCtx, batchIdToPrecision, ss, batchIdToCalSupply); err != nil {
		return err
	}

	basketStore, err := basketapi.NewStateStore(ormdb)
	if err != nil {
		return err
	}

	bBalanceItr, err := basketStore.BasketBalanceTable().List(ormCtx, basketapi.BasketBalancePrimaryKey{})
	if err != nil {
		return err
	}
	defer bBalanceItr.Close()

	for bBalanceItr.Next() {
		bBalance, err := bBalanceItr.Value()
		if err != nil {
			return err
		}
		batchId, ok := batchDenomToIdMap[bBalance.BatchDenom]
		if !ok {
			return fmt.Errorf("unknown credit batch %d in basket", batchId)
		}

		bb, err := math.NewNonNegativeDecFromString(bBalance.Balance)
		if err != nil {
			return err
		}

		if amount, ok := batchIdToCalSupply[batchId]; ok {
			result, err := math.SafeAddBalance(amount, bb)
			if err != nil {
				return err
			}
			batchIdToCalSupply[batchId] = result
		} else {
			return fmt.Errorf("unknown credit batch %d in basket", batchId)
		}
	}

	// verify calculated total amount of each credit batch matches the total supply
	if err := validateSupply(batchIdToCalSupply, batchIdToSupply); err != nil {
		return err
	}

	return nil
}

func validateMsg(m proto.Message) error {
	switch m.(type) {
	case *api.Class:
		msg := &Class{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}

		return msg.Validate()
	case *api.ClassIssuer:
		msg := &ClassIssuer{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}

		return msg.Validate()
	case *api.Project:
		msg := &Project{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}

		return msg.Validate()
	case *api.Batch:
		msg := &Batch{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.CreditType:
		msg := &CreditType{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	}

	return nil
}

func calculateSupply(ctx context.Context, batchIdToPrecision map[uint64]uint32, ss api.StateStore, batchIdToSupply map[uint64]math.Dec) error {
	bbItr, err := ss.BatchBalanceTable().List(ctx, api.BatchBalancePrimaryKey{})
	if err != nil {
		return err
	}
	defer bbItr.Close()

	for bbItr.Next() {
		tradable := math.NewDecFromInt64(0)
		retired := math.NewDecFromInt64(0)
		escrowed := math.NewDecFromInt64(0)

		balance, err := bbItr.Value()
		if err != nil {
			return err
		}

		precision, ok := batchIdToPrecision[balance.BatchKey]
		if !ok {
			return sdkerrors.ErrInvalidType.Wrapf("credit type not exist for %d batch", balance.BatchKey)
		}

		if balance.TradableAmount != "" {
			tradable, err = math.NewNonNegativeFixedDecFromString(balance.TradableAmount, precision)
			if err != nil {
				return err
			}
		}

		if balance.RetiredAmount != "" {
			retired, err = math.NewNonNegativeFixedDecFromString(balance.RetiredAmount, precision)
			if err != nil {
				return err
			}
		}

		if balance.EscrowedAmount != "" {
			escrowed, err = math.NewNonNegativeFixedDecFromString(balance.EscrowedAmount, precision)
			if err != nil {
				return err
			}
		}

		total, err := math.Add(tradable, retired)
		if err != nil {
			return err
		}

		total, err = math.Add(total, escrowed)
		if err != nil {
			return err
		}

		if supply, ok := batchIdToSupply[balance.BatchKey]; ok {
			result, err := math.SafeAddBalance(supply, total)
			if err != nil {
				return err
			}
			batchIdToSupply[balance.BatchKey] = result
		} else {
			batchIdToSupply[balance.BatchKey] = total
		}
	}

	return nil
}

func validateSupply(batchIdToSupplyCal, batchIdToSupply map[uint64]math.Dec) error {
	if len(batchIdToSupplyCal) == 0 && len(batchIdToSupply) > 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("batch supply was given but no balances were found")
	}
	if len(batchIdToSupply) == 0 && len(batchIdToSupplyCal) > 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("batch balances were given but no supplies were found")
	}
	for denom, cs := range batchIdToSupplyCal {
		if s, ok := batchIdToSupply[denom]; ok {
			if s.Cmp(cs) != math.EqualTo {
				return sdkerrors.ErrInvalidCoins.Wrapf("supply is incorrect for %d credit batch, expected %v, got %v", denom, s, cs)
			}
		} else {
			return sdkerrors.ErrNotFound.Wrapf("supply is not found for %d credit batch", denom)
		}
	}

	return nil
}

// MergeParamsIntoTarget merges params message into the ormjson.WriteTarget.
func MergeParamsIntoTarget(cdc codec.JSONCodec, message gogoproto.Message, target ormjson.WriteTarget) error {
	w, err := target.OpenWriter(protoreflect.FullName(gogoproto.MessageName(message)))
	if err != nil {
		return err
	}

	bz, err := cdc.MarshalJSON(message)
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	if err != nil {
		return err
	}

	return w.Close()
}

// Validate performs a basic validation of credit class
func (c Class) Validate() error {
	if len(c.Metadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrap("credit class metadata")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(c.Admin).String()); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if len(c.Id) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class id cannot be empty")
	}

	if err := ValidateClassId(c.Id); err != nil {
		return err
	}

	if len(c.CreditTypeAbbrev) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("must specify a credit type abbreviation")
	}

	return nil
}

// Validate performs a basic validation of credit class issuers
func (c ClassIssuer) Validate() error {
	if c.ClassKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class key cannot be zero")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(c.Issuer).String()); err != nil {
		return sdkerrors.Wrap(err, "issuer")
	}

	return nil
}

// Validate performs a basic validation of project
func (p Project) Validate() error {
	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(p.Admin).String()); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if p.ClassKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class key cannot be zero")
	}

	if err := ValidateJurisdiction(p.Jurisdiction); err != nil {
		return err
	}

	if len(p.Metadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrap("project metadata")
	}

	if err := ValidateProjectId(p.Id); err != nil {
		return err
	}

	return nil
}

// Validate performs a basic validation of credit batch
func (b Batch) Validate() error {
	if err := ValidateBatchDenom(b.Denom); err != nil {
		return err
	}

	if b.ProjectKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("project key cannot be zero")
	}

	if b.StartDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide a start date for the credit batch")
	}
	if b.EndDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide an end date for the credit batch")
	}
	if b.EndDate.Compare(*b.StartDate) != 1 {
		return sdkerrors.ErrInvalidRequest.Wrapf("the batch end date (%s) must be the same as or after the batch start date (%s)", b.EndDate.String(), b.StartDate.String())
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(b.Issuer).String()); err != nil {
		return sdkerrors.Wrap(err, "issuer")
	}

	return nil
}
