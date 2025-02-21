package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v2/x/staging/subspaces/types"
)

func (suite *KeeperTestsuite) TestQueryServer_Subspace() {
	tests := []struct {
		name        string
		store       func(ctx sdk.Context)
		request     *types.QuerySubspaceRequest
		shouldErr   bool
		expResponse *types.QuerySubspaceResponse
	}{
		{
			name:      "Invalid subspace id returns error",
			request:   types.NewQuerySubspaceRequest("123"),
			shouldErr: true,
		},
		{
			name:      "Not found subspace returns error",
			request:   types.NewQuerySubspaceRequest("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			shouldErr: true,
		},
		{
			name: "Found subspace is returned properly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			request:   types.NewQuerySubspaceRequest("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			shouldErr: false,
			expResponse: &types.QuerySubspaceResponse{
				Subspace: types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			response, err := suite.k.Subspace(sdk.WrapSDKContext(suite.ctx), test.request)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expResponse, response)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestQueryServer_Subspaces() {
	tests := []struct {
		name         string
		store        func(ctx sdk.Context)
		req          *types.QuerySubspacesRequest
		expSubspaces []types.Subspace
	}{

		{
			name: "Invalid pagination returns empty slice",
			req: types.NewQuerySubspacesRequest(&query.PageRequest{
				Limit:  1,
				Offset: 1,
			}),
			expSubspaces: nil,
		},
		{
			name: "Valid pagination returns result properly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			req: &types.QuerySubspacesRequest{},
			expSubspaces: []types.Subspace{
				types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			res, err := suite.k.Subspaces(sdk.WrapSDKContext(suite.ctx), test.req)
			suite.Require().NoError(err)
			suite.Require().Equal(test.expSubspaces, res.Subspaces)
		})
	}
}

func (suite *KeeperTestsuite) TestQueryServer_Admins() {
	tests := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryAdminsRequest
		shouldErr bool
		expAdmins []string
	}{
		{
			name:      "Invalid subspace id returns error",
			req:       types.NewQueryAdminsRequest("123", nil),
			shouldErr: true,
		},
		{
			name:      "Non existing subspace returns empty slice",
			req:       types.NewQueryAdminsRequest("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", nil),
			shouldErr: false,
			expAdmins: nil,
		},
		{
			name: "Requests pagination works properly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(suite.ctx, subspace.ID, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773", subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(suite.ctx, subspace.ID, "cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn", subspace.Owner)
				suite.Require().NoError(err)
			},
			req: types.NewQueryAdminsRequest(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				&query.PageRequest{
					Offset: 1,
					Limit:  1,
				},
			),
			expAdmins: []string{
				"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			res, err := suite.k.Admins(sdk.WrapSDKContext(suite.ctx), test.req)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expAdmins, res.Admins)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestQueryServer_RegisteredUsers() {
	tests := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryRegisteredUsersRequest
		shouldErr bool
		expUsers  []string
	}{
		{
			name:      "Invalid subspace id returns error",
			req:       types.NewQueryRegisteredUsersRequest("123", nil),
			shouldErr: true,
		},
		{
			name:      "Non existing subspace returns empty slice",
			req:       types.NewQueryRegisteredUsersRequest("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", nil),
			shouldErr: false,
			expUsers:  nil,
		},
		{
			name: "Requests pagination works properly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(suite.ctx, subspace.ID, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773", subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(suite.ctx, subspace.ID, "cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn", subspace.Owner)
				suite.Require().NoError(err)
			},
			req: types.NewQueryRegisteredUsersRequest(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				&query.PageRequest{
					Offset: 1,
					Limit:  1,
				},
			),
			expUsers: []string{
				"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			res, err := suite.k.RegisteredUsers(sdk.WrapSDKContext(suite.ctx), test.req)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expUsers, res.Users)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestQueryServer_BannedUsers() {
	tests := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryBannedUsersRequest
		shouldErr bool
		expUsers  []string
	}{
		{
			name:      "Invalid subspace id returns error",
			req:       types.NewQueryBannedUsersRequest("123", nil),
			shouldErr: true,
		},
		{
			name:      "Non existing subspace returns empty slice",
			req:       types.NewQueryBannedUsersRequest("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", nil),
			shouldErr: false,
			expUsers:  nil,
		},
		{
			name: "Requests pagination works properly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(suite.ctx, subspace.ID, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773", subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(suite.ctx, subspace.ID, "cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn", subspace.Owner)
				suite.Require().NoError(err)
			},
			req: types.NewQueryBannedUsersRequest(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				&query.PageRequest{
					Offset: 1,
					Limit:  1,
				},
			),
			expUsers: []string{
				"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			res, err := suite.k.BannedUsers(sdk.WrapSDKContext(suite.ctx), test.req)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expUsers, res.Users)
			}
		})
	}
}
