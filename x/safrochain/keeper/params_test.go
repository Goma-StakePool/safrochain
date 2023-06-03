package keeper_test

import (
	"testing"

	testkeeper "github.com/Goma-StakePool/safrochain/testutil/keeper"
	"github.com/Goma-StakePool/safrochain/x/safrochain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SafrochainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
