package kiichain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"kiichain/x/kiichain/keeper"
	"kiichain/x/kiichain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the post
	for _, elem := range genState.PostList {
		k.SetPost(ctx, elem)
	}

	// Set post count
	k.SetPostCount(ctx, genState.PostCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PostList = k.GetAllPost(ctx)
	genesis.PostCount = k.GetPostCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
