package keeper

import (
	"github.com/chainstock-project/blockchain/x/blockchain/types"
)

var _ types.QueryServer = Keeper{}
