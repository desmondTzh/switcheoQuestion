package desmondtzh_test

import (
	"testing"

	keepertest "desmondTZH/testutil/keeper"
	"desmondTZH/testutil/nullify"
	desmondtzh "desmondTZH/x/desmondtzh/module"
	"desmondTZH/x/desmondtzh/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DesmondtzhKeeper(t)
	desmondtzh.InitGenesis(ctx, k, genesisState)
	got := desmondtzh.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
