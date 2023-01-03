package bc_test

import (
	"testing"

	keepertest "bc/testutil/keeper"
	"bc/testutil/nullify"
	"bc/x/bc"
	"bc/x/bc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BcKeeper(t)
	bc.InitGenesis(ctx, *k, genesisState)
	got := bc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
