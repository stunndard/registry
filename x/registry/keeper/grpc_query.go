package keeper

import (
	"github.com/stunndard/registry/x/registry/types"
)

var _ types.QueryServer = Keeper{}
