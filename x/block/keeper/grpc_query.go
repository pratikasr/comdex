package keeper

import (
	"github.com/comdex-official/comdex/x/block/types"
)

var _ types.QueryServer = Keeper{}
