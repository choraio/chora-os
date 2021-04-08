package keeper

import (
	"github.com/choraio/chora/x/ecodex/types"
)

var _ types.QueryServer = Keeper{}
