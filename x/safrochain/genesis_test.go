package safrochain_test

import (
	"testing"

	keepertest "github.com/Goma-StakePool/safrochain/testutil/keeper"
	"github.com/Goma-StakePool/safrochain/testutil/nullify"
	"github.com/Goma-StakePool/safrochain/x/safrochain"
	"github.com/Goma-StakePool/safrochain/x/safrochain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SafrochainKeeper(t)
	safrochain.InitGenesis(ctx, *k, genesisState)
	got := safrochain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
