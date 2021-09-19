package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
)

func TestStockDataMsgServerCreate(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	srv := NewMsgServerImpl(*keeper)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		idx := fmt.Sprintf("%d", i)
		expected := &types.MsgCreateStockData{Creator: creator, Index: idx}
		_, err := srv.CreateStockData(wctx, expected)
		require.NoError(t, err)
		rst, found := keeper.GetStockData(ctx, expected.Index)
		require.True(t, found)
		assert.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestStockDataMsgServerUpdate(t *testing.T) {
	creator := "A"
	index := "any"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateStockData
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateStockData{Creator: creator, Index: index},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateStockData{Creator: "B", Index: index},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgUpdateStockData{Creator: creator, Index: "missing"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			keeper, ctx := setupKeeper(t)
			srv := NewMsgServerImpl(*keeper)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateStockData{Creator: creator, Index: index}
			_, err := srv.CreateStockData(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateStockData(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := keeper.GetStockData(ctx, expected.Index)
				require.True(t, found)
				assert.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestStockDataMsgServerDelete(t *testing.T) {
	creator := "A"
	index := "any"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteStockData
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteStockData{Creator: creator, Index: index},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteStockData{Creator: "B", Index: index},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteStockData{Creator: creator, Index: "missing"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			keeper, ctx := setupKeeper(t)
			srv := NewMsgServerImpl(*keeper)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateStockData(wctx, &types.MsgCreateStockData{Creator: creator, Index: index})
			require.NoError(t, err)
			_, err = srv.DeleteStockData(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := keeper.GetStockData(ctx, tc.request.Index)
				require.False(t, found)
			}
		})
	}
}
