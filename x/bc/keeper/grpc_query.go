package keeper

import (
	"bc/x/bc/types"
)

var _ types.QueryServer = Keeper{}
