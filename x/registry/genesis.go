package registry

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stunndard/registry/x/registry/keeper"
	"github.com/stunndard/registry/x/registry/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the name
	for _, elem := range genState.NameList {
		k.SetName(ctx, *elem)
	}

	// Set name count
	k.SetNameCount(ctx, int64(len(genState.NameList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all name
	nameList := k.GetAllName(ctx)
	for _, elem := range nameList {
		genesis.NameList = append(genesis.NameList, &elem)
	}

	return genesis
}
