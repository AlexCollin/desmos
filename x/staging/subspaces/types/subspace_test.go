package types_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/x/staging/subspaces/types"
)

func TestSubspace_WithName(t *testing.T) {
	sub := types.NewSubspace(
		"123",
		"name",
		"",
		"",
		types.SubspaceTypeOpen,
		time.Unix(1, 2),
	).WithName("sub")
	require.Equal(t, "sub", sub.Name)
}

func TestSubspace_WithOwner(t *testing.T) {
	sub := types.NewSubspace(
		"123",
		"name",
		"",
		"",
		types.SubspaceTypeOpen,
		time.Unix(1, 2),
	).WithOwner("owner")
	require.Equal(t, "owner", sub.Owner)
}

func TestSubspace_WithSubspaceType(t *testing.T) {
	sub := types.NewSubspace(
		"123",
		"name",
		"",
		"",
		types.SubspaceTypeOpen,
		time.Unix(1, 2),
	).WithSubspaceType(types.SubspaceTypeClosed)
	require.Equal(t, types.SubspaceTypeClosed, sub.Type)
}

func TestSubspace_Validate(t *testing.T) {
	date := time.Date(2050, 01, 01, 15, 15, 00, 000, time.UTC)
	tests := []struct {
		name     string
		subspace types.Subspace
		expError bool
	}{
		{
			name: "Invalid ID returns error",
			subspace: types.NewSubspace(
				"123",
				"",
				"",
				"",
				types.SubspaceTypeOpen,
				time.Time{},
			),
			expError: true,
		},
		{
			name: "Invalid name returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"",
				"",
				"",
				types.SubspaceTypeOpen,
				time.Time{},
			),
			expError: true,
		},
		{
			name: "Invalid owner returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"",
				"",
				types.SubspaceTypeOpen,
				time.Time{},
			),
			expError: true,
		},
		{
			name: "Invalid creator returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"",
				types.SubspaceTypeOpen,
				time.Time{},
			),
			expError: true,
		},
		{
			name: "Invalid creation time returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				types.SubspaceTypeOpen,
				time.Time{},
			),
			expError: true,
		},
		{
			name: "Valid subspace returns no error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				types.SubspaceTypeOpen,
				date,
			),
			expError: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.subspace.Validate()
			if test.expError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_IsValidSubspaceType(t *testing.T) {
	tests := []struct {
		name     string
		subType  types.SubspaceType
		expValid bool
	}{
		{
			name:     "valid open type returns true",
			subType:  types.SubspaceTypeOpen,
			expValid: true,
		},
		{
			name:     "valid close type returns true",
			subType:  types.SubspaceTypeClosed,
			expValid: true,
		},
		{
			name:     "invalid type returns false",
			subType:  types.SubspaceTypeUnspecified,
			expValid: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expValid, types.IsValidSubspaceType(test.subType))
		})
	}
}

func Test_NormalizeSubspaceType(t *testing.T) {
	tests := []struct {
		name       string
		subType    string
		expSubType string
	}{
		{
			name:       "Valid Open subspace Type",
			subType:    "open",
			expSubType: types.SubspaceTypeOpen.String(),
		},
		{
			name:       "Valid Close subspace type",
			subType:    "Close",
			expSubType: types.SubspaceTypeClosed.String(),
		},
		{
			name:       "Invalid subspace type",
			subType:    "Invalid",
			expSubType: "Invalid",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {

			subspaceType := types.NormalizeSubspaceType(test.subType)
			require.Equal(t, test.expSubType, subspaceType)
		})
	}
}

func Test_SubspaceTypeFromString(t *testing.T) {
	tests := []struct {
		name       string
		subType    string
		expSubType types.SubspaceType
		expError   error
	}{
		{
			name:       "Invalid subspace type",
			subType:    "invalid",
			expSubType: types.SubspaceTypeUnspecified,
			expError:   fmt.Errorf("'invalid' is not a valid subspace type"),
		},
		{
			name:       "Valid subspace type",
			subType:    types.SubspaceTypeOpen.String(),
			expSubType: types.SubspaceTypeOpen,
			expError:   nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			res, err := types.SubspaceTypeFromString(test.subType)
			require.Equal(t, test.expError, err)
			require.Equal(t, test.expSubType, res)
		})
	}
}
