package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/incubus-network/fury/x/furydist/types"
)

// Keeper keeper for the cdp module
type Keeper struct {
	key           storetypes.StoreKey
	cdc           codec.BinaryCodec
	paramSubspace paramtypes.Subspace
	bankKeeper    types.BankKeeper
	distKeeper    types.DistKeeper
	accountKeeper types.AccountKeeper

	blacklistedAddrs map[string]bool
}

// NewKeeper creates a new keeper
func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, paramstore paramtypes.Subspace, bk types.BankKeeper, ak types.AccountKeeper,
	dk types.DistKeeper, blacklistedAddrs map[string]bool,
) Keeper {
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		key:              key,
		cdc:              cdc,
		paramSubspace:    paramstore,
		bankKeeper:       bk,
		distKeeper:       dk,
		accountKeeper:    ak,
		blacklistedAddrs: blacklistedAddrs,
	}
}

// GetPreviousBlockTime get the blocktime for the previous block
func (k Keeper) GetPreviousBlockTime(ctx sdk.Context) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousBlockTimeKey)
	b := store.Get(types.PreviousBlockTimeKey)
	if b == nil {
		return time.Time{}, false
	}
	if err := blockTime.UnmarshalBinary(b); err != nil {
		return time.Time{}, false
	}
	return blockTime, true
}

// SetPreviousBlockTime set the time of the previous block
func (k Keeper) SetPreviousBlockTime(ctx sdk.Context, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousBlockTimeKey)
	b, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set(types.PreviousBlockTimeKey, b)
}
