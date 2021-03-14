package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stunndard/registry/x/registry/types"
)

// GetNameCount get the total number of name
func (k Keeper) GetNameCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameCountKey))
	byteKey := types.KeyPrefix(types.NameCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetNameCount set the total number of name
func (k Keeper) SetNameCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameCountKey))
	byteKey := types.KeyPrefix(types.NameCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateName creates a name with a new id and update the count
func (k Keeper) CreateName(ctx sdk.Context, msg types.MsgBuyName) {
	// Create the name
	count := k.GetNameCount(ctx)
	name := types.Name{
		Creator: msg.Creator,
		Id:      strconv.FormatInt(count, 10),
		Owner:   msg.Creator,
		Name:    msg.Name,
		Price:   msg.Bid,
		Onsale:  msg.Onsale,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	key := types.KeyPrefix(types.NameKey + name.Name)
	value := k.cdc.MustMarshalBinaryBare(&name)
	store.Set(key, value)

	// Update name count
	k.SetNameCount(ctx, count+1)
}

// SetName set a specific name in the store
func (k Keeper) SetName(ctx sdk.Context, name types.Name) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	b := k.cdc.MustMarshalBinaryBare(&name)
	store.Set(types.KeyPrefix(types.NameKey+name.Name), b)
}

// GetName returns a name
func (k Keeper) GetName(ctx sdk.Context, key string) types.Name {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	var name types.Name
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.NameKey+key)), &name)
	return name
}

// HasName checks if the name exists
func (k Keeper) HasName(ctx sdk.Context, name string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	return store.Has(types.KeyPrefix(types.NameKey + name))
}

// GetNameOwner returns the owner of the name
func (k Keeper) GetNameOwner(ctx sdk.Context, name string) string {
	return k.GetName(ctx, name).Owner
}

// GetNameCreator returns the creator of the name
func (k Keeper) GetNameCreator(ctx sdk.Context, name string) string {
	return k.GetName(ctx, name).Creator
}

// DeleteName deletes a name
func (k Keeper) DeleteName(ctx sdk.Context, name string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	store.Delete(types.KeyPrefix(types.NameKey + name))
}

// GetAllName returns all name
func (k Keeper) GetAllName(ctx sdk.Context) (msgs []types.Name) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.NameKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Name
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

// GetPrice returns the price of a name
func (k Keeper) GetPrice(ctx sdk.Context, key string) sdk.Coins {
	name := k.GetName(ctx, key)
	return name.Price
}
