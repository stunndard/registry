package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stunndard/registry/x/registry/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		CoinKeeper keeper.Keeper
		cdc        codec.Marshaler
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
	}
)

func NewKeeper(coinKeeper keeper.Keeper, cdc codec.Marshaler, storeKey, memKey sdk.StoreKey) *Keeper {
	return &Keeper{
		CoinKeeper: coinKeeper,
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
