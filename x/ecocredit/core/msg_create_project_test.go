package core

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgCreateProject(t *testing.T) {
	t.Parallel()
	admin := testutil.GenAddress()

	testCases := []struct {
		name   string
		src    MsgCreateProject
		expErr bool
	}{
		{
			"valid msg without reference id",
			MsgCreateProject{
				Admin:        admin,
				ClassId:      "A00",
				Metadata:     "hello",
				Jurisdiction: "AB-CDE FG1 345",
			},
			false,
		},
		{
			"invalid admin",
			MsgCreateProject{
				Admin:        "invalid address",
				ClassId:      "A00",
				Metadata:     "hello",
				Jurisdiction: "AB-CDE FG1 345",
			},
			true,
		},
		{
			"invalid class id",
			MsgCreateProject{
				Admin:        admin,
				ClassId:      "ABCD",
				Metadata:     "hello",
				Jurisdiction: "AB-CDE FG1 345",
			},
			true,
		},
		{
			"invalid project jurisdiction",
			MsgCreateProject{
				Admin:        admin,
				ClassId:      "A01",
				Metadata:     "hello",
				Jurisdiction: "abcd",
			},
			true,
		},
		{
			"invalid: metadata is too large",
			MsgCreateProject{
				Admin:        admin,
				ClassId:      "A01",
				Metadata:     strings.Repeat("x", 288),
				Jurisdiction: "AB-CDE FG1 345",
			},
			true,
		},
		{
			"invalid: reference id is too large",
			MsgCreateProject{
				Admin:        admin,
				ClassId:      "A01",
				Metadata:     "metadata",
				Jurisdiction: "AB-CDE FG1 345",
				ReferenceId:  strings.Repeat("x", MaxReferenceIdLength+1),
			},
			true,
		},
		{
			"valid: with reference id",
			MsgCreateProject{
				Admin:        admin,
				ClassId:      "A01",
				Metadata:     "metadata",
				Jurisdiction: "AB-CDE FG1 345",
				ReferenceId:  strings.Repeat("x", 10),
			},
			false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
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
