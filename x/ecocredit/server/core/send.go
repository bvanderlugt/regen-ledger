package core

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Send sends credits to a recipient.
// Send also retires credits if the amount to retire is specified in the request.
func (k Keeper) Send(ctx context.Context, req *core.MsgSend) (*core.MsgSendResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	sender, _ := sdk.AccAddressFromBech32(req.Sender)
	recipient, _ := sdk.AccAddressFromBech32(req.Recipient)

	for _, credit := range req.Credits {
		err := k.sendEcocredits(ctx, credit, recipient, sender)
		if err != nil {
			return nil, err
		}
		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventTransfer{
			Sender:         req.Sender,
			Recipient:      req.Recipient,
			BatchDenom:     credit.BatchDenom,
			TradableAmount: credit.TradableAmount,
			RetiredAmount:  credit.RetiredAmount,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgSend credit iteration")
	}
	return &core.MsgSendResponse{}, nil
}

func (k Keeper) sendEcocredits(ctx context.Context, credit *core.MsgSend_SendCredits, to, from sdk.AccAddress) error {
	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, credit.BatchDenom)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("could not get batch with denom %s: %s", credit.BatchDenom, err.Error())
	}
	creditType, err := utils.GetCreditTypeFromBatchDenom(ctx, k.stateStore, batch.Denom)
	if err != nil {
		return err
	}
	precision := creditType.Precision

	batchSupply, err := k.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
	if err != nil {
		return err
	}
	fromBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, from, batch.Key)
	if err != nil {
		if err == ormerrors.NotFound {
			return ecocredit.ErrInsufficientCredits.Wrapf("you do not have any credits from batch %s", batch.Denom)
		}
		return err
	}

	toBalance, err := k.stateStore.BatchBalanceTable().Get(ctx, to, batch.Key)
	if err != nil {
		if err == ormerrors.NotFound {
			toBalance = &api.BatchBalance{
				BatchKey:       batch.Key,
				Address:        to,
				TradableAmount: "0",
				RetiredAmount:  "0",
			}
		} else {
			return err
		}
	}
	decs, err := utils.GetNonNegativeFixedDecs(precision, toBalance.TradableAmount, toBalance.RetiredAmount, fromBalance.TradableAmount, fromBalance.RetiredAmount, credit.TradableAmount, credit.RetiredAmount, batchSupply.TradableAmount, batchSupply.RetiredAmount)
	if err != nil {
		return err
	}
	toTradableBalance, toRetiredBalance,
		fromTradableBalance, fromRetiredBalance,
		sendAmtTradable, sendAmtRetired,
		batchSupplyTradable, batchSupplyRetired := decs[0], decs[1], decs[2], decs[3], decs[4], decs[5], decs[6], decs[7]

	if !sendAmtTradable.IsZero() {
		fromTradableBalance, err = math.SafeSubBalance(fromTradableBalance, sendAmtTradable)
		if err != nil {
			return err
		}
		toTradableBalance, err = toTradableBalance.Add(sendAmtTradable)
		if err != nil {
			return err
		}
	}

	didRetire := false
	if !sendAmtRetired.IsZero() {
		didRetire = true
		fromTradableBalance, err = math.SafeSubBalance(fromTradableBalance, sendAmtRetired)
		if err != nil {
			return err
		}
		toRetiredBalance, err = toRetiredBalance.Add(sendAmtRetired)
		if err != nil {
			return err
		}
		batchSupplyRetired, err = batchSupplyRetired.Add(sendAmtRetired)
		if err != nil {
			return err
		}
		batchSupplyTradable, err = batchSupplyTradable.Sub(sendAmtRetired)
		if err != nil {
			return err
		}
	}
	// update the "to" balance
	if err := k.stateStore.BatchBalanceTable().Save(ctx, &api.BatchBalance{
		BatchKey:       batch.Key,
		Address:        to,
		TradableAmount: toTradableBalance.String(),
		RetiredAmount:  toRetiredBalance.String(),
	}); err != nil {
		return err
	}

	// update the "from" balance
	if err := k.stateStore.BatchBalanceTable().Update(ctx, &api.BatchBalance{
		BatchKey:       batch.Key,
		Address:        from,
		TradableAmount: fromTradableBalance.String(),
		RetiredAmount:  fromRetiredBalance.String(),
	}); err != nil {
		return err
	}
	// update the "retired" supply only if credits were retired
	if didRetire {
		if err := k.stateStore.BatchSupplyTable().Update(ctx, &api.BatchSupply{
			BatchKey:        batch.Key,
			TradableAmount:  batchSupplyTradable.String(),
			RetiredAmount:   batchSupplyRetired.String(),
			CancelledAmount: batchSupply.CancelledAmount,
		}); err != nil {
			return err
		}
		if err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&core.EventRetire{
			Owner:        to.String(),
			BatchDenom:   credit.BatchDenom,
			Amount:       sendAmtRetired.String(),
			Jurisdiction: credit.RetirementJurisdiction,
		}); err != nil {
			return err
		}
	}
	return nil
}
