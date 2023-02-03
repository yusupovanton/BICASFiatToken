package bicasfiattoken_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/yusupovanton/BICASFiatToken/testutil/keeper"
	"github.com/yusupovanton/BICASFiatToken/testutil/nullify"
	"github.com/yusupovanton/BICASFiatToken/x/bicasfiattoken"
	"github.com/yusupovanton/BICASFiatToken/x/bicasfiattoken/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BicasfiattokenKeeper(t)
	bicasfiattoken.InitGenesis(ctx, *k, genesisState)
	got := bicasfiattoken.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
