package furydist

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/incubus-network/fury/x/furydist/keeper"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	err := k.MintPeriodInflation(ctx)
	if err != nil {
		panic(err)
	}
}
