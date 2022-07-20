package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "kwil/testutil/keeper"
	"kwil/testutil/nullify"
	"kwil/x/kwil/keeper"
	"kwil/x/kwil/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNQueryids(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Queryids {
	items := make([]types.Queryids, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetQueryids(ctx, items[i])
	}
	return items
}

func TestQueryidsGet(t *testing.T) {
	keeper, ctx := keepertest.KwilKeeper(t)
	items := createNQueryids(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetQueryids(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestQueryidsRemove(t *testing.T) {
	keeper, ctx := keepertest.KwilKeeper(t)
	items := createNQueryids(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveQueryids(ctx,
			item.Index,
		)
		_, found := keeper.GetQueryids(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestQueryidsGetAll(t *testing.T) {
	keeper, ctx := keepertest.KwilKeeper(t)
	items := createNQueryids(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllQueryids(ctx)),
	)
}
