package ratnet_test

import (
	"testing"

	keepertest "github.com/GuillaumeDonnet/RatNet/testutil/keeper"
	"github.com/GuillaumeDonnet/RatNet/testutil/nullify"
	ratnet "github.com/GuillaumeDonnet/RatNet/x/ratnet/module"
	"github.com/GuillaumeDonnet/RatNet/x/ratnet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RatnetKeeper(t)
	ratnet.InitGenesis(ctx, k, genesisState)
	got := ratnet.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
