package basketclient

import (
	"fmt"
	"strings"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

const (
	FlagDisableAutoRetire      = "disable-auto-retire"
	FlagCreditTypeAbbreviation = "credit-type-abbreviation"
	FlagAllowedClasses         = "allowed-classes"
	FlagMinimumStartDate       = "minimum-start-date"
	FlagStartDateWindow        = "start-date-window"
	FlagBasketFee              = "basket-fee"
	FlagDenomDescription       = "description"
	FlagRetirementJurisdiction = "retirement-jurisdiction"
	FlagRetireOnTake           = "retire-on-take"
)

func TxCreateBasketCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-basket [name]",
		Short: "Creates a bank denom that wraps credits",
		Long: strings.TrimSpace(`Creates a bank denom that wraps credits

Parameters:
		name: the name used to create a bank denom for this basket token.

Flags:
		disable-auto-retire: disables the auto-retirement of credits upon taking credits
			from the basket. The credits will be auto-retired if disable_auto_retire is
			false unless the credits were previously put into the basket by the address
			picking them from the basket, in which case they will remain tradable.
		credit-type-abbreviation: filters against credits from this credit type abbreviation (e.g. "BIO").
		allowed_classes: comma separated (no spaces) list of credit classes allowed to be put in
			the basket (e.g. "C01,C02").
		min-start-date: the earliest start date for batches of credits allowed into the basket.
		start-date-window: the duration of time (in seconds) measured into the past which sets a
			cutoff for batch start dates when adding new credits to the basket.
		basket-fee: the fee that the curator will pay to create the basket. It must be >= the
			required Params.basket_creation_fee. We include the fee explicitly here so that the
			curator explicitly acknowledges paying this fee and is not surprised to learn that the
			paid a big fee and didn't know beforehand.
		description: the description to be used in the basket coin's bank denom metadata.`),
		Example: `
		$regen tx ecocredit create-basket HEAED
			--from regen...
			--credit-type-abbreviation=FOO
			--allowed_classes="class1,class2"
			--basket-fee=100regen
			--description="any description"
		`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			disableAutoRetire, err := cmd.Flags().GetBool(FlagDisableAutoRetire)
			if err != nil {
				return err
			}

			creditTypeName, err := cmd.Flags().GetString(FlagCreditTypeAbbreviation)
			if err != nil {
				return err
			}

			allowedClasses, err := cmd.Flags().GetStringSlice(FlagAllowedClasses)
			if err != nil {
				return err
			}
			for i := range allowedClasses {
				allowedClasses[i] = strings.TrimSpace(allowedClasses[i])
			}

			minStartDateString, err := cmd.Flags().GetString(FlagMinimumStartDate)
			if err != nil {
				return err
			}
			startDateWindow, err := cmd.Flags().GetUint64(FlagStartDateWindow)
			if err != nil {
				return err
			}

			denomDescription, err := cmd.Flags().GetString(FlagDenomDescription)
			if err != nil {
				return err
			}

			if minStartDateString != "" && startDateWindow != 0 {
				return fmt.Errorf("both %s and %s cannot be set", FlagStartDateWindow, FlagMinimumStartDate)
			}

			var dateCriteria *basket.DateCriteria

			if minStartDateString != "" {
				minStartDateTime, err := regentypes.ParseDate("min-start-date", minStartDateString)
				if err != nil {
					return err
				}
				minStartDate, err := types.TimestampProto(minStartDateTime)
				if err != nil {
					return fmt.Errorf("failed to parse min_start_date: %w", err)
				}
				dateCriteria = &basket.DateCriteria{MinStartDate: minStartDate}
			}

			if startDateWindow != 0 {
				startDateWindowDuration := time.Duration(startDateWindow)
				startDateWindow := types.DurationProto(startDateWindowDuration)
				dateCriteria = &basket.DateCriteria{StartDateWindow: startDateWindow}
			}

			fee := sdk.Coins{}
			feeString, err := cmd.Flags().GetString(FlagBasketFee)
			if err != nil {
				return err
			}
			if feeString != "" {
				fee, err = sdk.ParseCoinsNormalized(feeString)
				if err != nil {
					return fmt.Errorf("failed to parse basket_fee: %w", err)
				}
			}

			msg := basket.MsgCreate{
				Curator:           clientCtx.FromAddress.String(),
				Name:              args[0],
				Description:       denomDescription,
				DisableAutoRetire: disableAutoRetire,
				CreditTypeAbbrev:  creditTypeName,
				AllowedClasses:    allowedClasses,
				DateCriteria:      dateCriteria,
				Fee:               fee,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	// command flags
	cmd.Flags().Bool(FlagDisableAutoRetire, false, "dictates whether credits will be auto-retired upon taking")
	cmd.Flags().String(FlagCreditTypeAbbreviation, "", "filters against credits from this credit type abbreviation (e.g. \"C\")")
	cmd.Flags().StringSlice(FlagAllowedClasses, []string{}, "comma separated (no spaces) list of credit classes allowed to be put in the basket (e.g. \"C01,C02\")")
	cmd.Flags().String(FlagMinimumStartDate, "", "the earliest start date for batches of credits allowed into the basket (e.g. \"2012-01-01\")")
	cmd.Flags().Uint64(FlagStartDateWindow, 0, "sets a cutoff for batch start dates when adding new credits to the basket (e.g. 1325404800)")
	cmd.Flags().String(FlagBasketFee, "", "the fee that the curator will pay to create the basket (e.g. \"20regen\")")
	cmd.Flags().String(FlagDenomDescription, "", "the description to be used in the bank denom metadata.")

	// required flags
	cmd.MarkFlagRequired(FlagCreditTypeAbbreviation)
	cmd.MarkFlagRequired(FlagAllowedClasses)

	return txFlags(cmd)
}

func TxPutInBasketCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put-in-basket [basket_denom] [credits]",
		Short: "add credits to the basket",
		Long: strings.TrimSpace(`add credits to the basket.
Parameters:
		basket_denom: basket identifier
		credits: path to JSON file containing credits to put in the basket
Flags:
		from: account address of the owner
		`),
		Example: `
regen tx ecocredit put-in-basket eco.uC.NCT credits.json

Where the credits.json file contains:

[
	{
		"batch_denom": "C01-001-20210101-20220101-001",
		"amount": "10"
	},
	{
		"batch_denom": "C01-001-20210101-20220101-001",
		"amount": "10.5"
	}
]
		`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			credits, err := parseBasketCredits(args[1])
			if err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("failed to parse json: %s", err)
			}

			msg := basket.MsgPut{
				Owner:       clientCtx.FromAddress.String(),
				BasketDenom: args[0],
				Credits:     credits,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return txFlags(cmd)
}

func TxTakeFromBasketCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "take-from-basket [basket_denom] [amount]",
		Short: "Takes credits from a basket",
		Long: strings.TrimSpace(`takes credits from a basket starting from the oldest credits first.
Parameters:
		basket_denom: denom identifying basket from which we redeem credits.
		amount: amount is a positive integer number of basket tokens to convert into credits.
Flags:
		from: account address of the owner of the basket.
		retirement-jurisdiction: retirement jurisdiction is the optional retirement jurisdiction for the credits
				which will be used only if --retire-on-take flag is true.
		retire-on-take: retire on take is a boolean that dictates whether the ecocredits
		                received in exchange for the basket tokens will be received as
		                retired or tradable credits.
		
		`),
		Example: `
regen tx ecocredit take-from-basket eco.uC.NCT 1000
regen tx ecocredit take-from-basket eco.uC.NCT 1000 --retire-on-take true --retirement-jurisdiction AQ
		`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			retirementJurisdiction, err := cmd.Flags().GetString(FlagRetirementJurisdiction)
			if err != nil {
				return err
			}

			retireOnTake, err := cmd.Flags().GetBool(FlagRetireOnTake)
			if err != nil {
				return err
			}

			msg := basket.MsgTake{
				Owner:                  clientCtx.FromAddress.String(),
				BasketDenom:            args[0],
				Amount:                 args[1],
				RetirementJurisdiction: retirementJurisdiction,
				RetireOnTake:           retireOnTake,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String(FlagRetirementJurisdiction, "", "jurisdiction for the credits which will be used only if --retire-on-take flag is true")
	cmd.Flags().Bool(FlagRetireOnTake, false, "dictates whether the ecocredits received in exchange for the basket tokens will be received as retired or tradable credits")

	return txFlags(cmd)
}
