package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

const MaxReferenceIdLength = 32

var _ legacytx.LegacyMsg = &MsgCreateProject{}

// Route implements the LegacyMsg interface.
func (m MsgCreateProject) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCreateProject) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateProject) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateProject) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("admin")
	}

	if err := ValidateClassId(m.ClassId); err != nil {
		return err
	}

	if len(m.Metadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrap("create project metadata")
	}

	if err := ValidateJurisdiction(m.Jurisdiction); err != nil {
		return err
	}

	if m.ReferenceId != "" && len(m.ReferenceId) > MaxReferenceIdLength {
		return ecocredit.ErrMaxLimit.Wrap("reference id")
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateProject.
func (m *MsgCreateProject) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}
