package keeper

import (
	"kwil/x/kwil/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDdl set a specific ddl in the store from its index
func (k Keeper) SetDdl(ctx sdk.Context, ddl types.Ddl) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DdlKeyPrefix))

	// We need to check to ensure that this is one position past the current index.
	// If it is the first, then this check does not need to take place
	if ddl.Position > 0 {
		k.GetDdl
	}

	b := k.cdc.MustMarshal(&ddl)
	store.Set(types.DdlKey(
		ddl.Index,
	), b)
}

// GetDdl returns a ddl from its index
func (k Keeper) GetDdl(
	ctx sdk.Context,
	index string,

) (val types.Ddl, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DdlKeyPrefix))

	b := store.Get(types.DdlKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDdl removes a ddl from the store
func (k Keeper) RemoveDdl(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DdlKeyPrefix))
	store.Delete(types.DdlKey(
		index,
	))
}

// GetAllDdl returns all ddl
func (k Keeper) GetAllDdl(ctx sdk.Context) (list []types.Ddl) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DdlKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Ddl
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
