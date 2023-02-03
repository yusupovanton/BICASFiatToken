package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/yusupovanton/BICASFiatToken/testutil/keeper"
	"github.com/yusupovanton/BICASFiatToken/x/bicasfiattoken/keeper"
	"github.com/yusupovanton/BICASFiatToken/x/bicasfiattoken/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BicasfiattokenKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
