package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "kwil/testutil/keeper"
	"kwil/x/kwil/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.KwilKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
