package keeper

import (
	"github.com/Goma-StakePool/safrochain/x/safrochain/types"
)

var _ types.QueryServer = Keeper{}
