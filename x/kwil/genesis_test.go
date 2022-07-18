package kwil_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "kwil/testutil/keeper"
	"kwil/testutil/nullify"
	"kwil/x/kwil"
	"kwil/x/kwil/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DatabasesList: []types.Databases{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KwilKeeper(t)
	kwil.InitGenesis(ctx, *k, genesisState)
	got := kwil.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DatabasesList, got.DatabasesList)
	// this line is used by starport scaffolding # genesis/test/assert
}
