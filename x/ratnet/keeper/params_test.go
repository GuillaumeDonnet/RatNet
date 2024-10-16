package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/GuillaumeDonnet/RatNet/testutil/keeper"
	"github.com/GuillaumeDonnet/RatNet/x/ratnet/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.RatnetKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
