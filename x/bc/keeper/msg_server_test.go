package keeper_test

import (
	"context"
	"testing"

	keepertest "bc/testutil/keeper"
	"bc/x/bc/keeper"
	"bc/x/bc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BcKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
