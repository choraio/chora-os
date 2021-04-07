package keeper

import (
	"github.com/choraio/chora/x/chora/types"
)

var _ types.QueryServer = Keeper{}
