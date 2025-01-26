package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "desmondtzh/testutil/keeper"
	"desmondtzh/x/desmondtzh/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DesmondtzhKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
