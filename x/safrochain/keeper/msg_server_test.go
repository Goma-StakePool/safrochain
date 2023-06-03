package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Goma-StakePool/safrochain/testutil/keeper"
	"github.com/Goma-StakePool/safrochain/x/safrochain/keeper"
	"github.com/Goma-StakePool/safrochain/x/safrochain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SafrochainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
