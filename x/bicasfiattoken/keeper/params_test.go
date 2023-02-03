package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/yusupovanton/BICASFiatToken/testutil/keeper"
	"github.com/yusupovanton/BICASFiatToken/x/bicasfiattoken/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BicasfiattokenKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
