package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgRetire(t *testing.T) {
	t.Parallel()

	addr1 := testutil.GenAddress()

	tests := map[string]struct {
		src    MsgRetire
		expErr bool
	}{
		"valid msg": {
			src: MsgRetire{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: false,
		},
		"invalid msg without holder": {
			src: MsgRetire{
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg with wrong holder address": {
			src: MsgRetire{
				Owner: "wrong owner",
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without credits": {
			src: MsgRetire{
				Owner:        addr1,
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without Credits.BatchDenom": {
			src: MsgRetire{
				Owner: addr1,
				Credits: []*Credits{
					{
						Amount: "10",
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without Credits.Amount": {
			src: MsgRetire{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg with wrong Credits.Amount": {
			src: MsgRetire{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "abc",
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without jurisdiction": {
			src: MsgRetire{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg with wrong jurisdiction": {
			src: MsgRetire{
				Owner: addr1,
				Credits: []*Credits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Jurisdiction: "wrongJurisdiction",
			},
			expErr: true,
		},
	}

	for msg, test := range tests {
		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			err := test.src.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
