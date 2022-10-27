package v2

import (
	"fmt"
	v2types "github.com/comdex-official/comdex/x/block/migrations/v2/types"
	"github.com/comdex-official/comdex/x/block/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuftypes "github.com/gogo/protobuf/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	fmt.Println("MigrateStore")

	return migrateValues(store, cdc)
}

func migrateValues(store sdk.KVStore, cdc codec.BinaryCodec) error {
	fmt.Println(" migrateValues")
	var (
		iter = sdk.KVStorePrefixIterator(store, types.AssetKeyPrefix)
	)

	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)
	var assets []v2types.Asset
	for ; iter.Valid(); iter.Next() {
		var asset v2types.Asset
		cdc.MustUnmarshal(iter.Value(), &asset)
		assets = append(assets, asset)
	}
	counterKey := types.NewAssetIDKey
	for _, v := range assets {
		newVal, Key := migrateValue(v)
		store.Delete(Key)
		value := cdc.MustMarshal(&newVal)
		idValue := cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: v.Id,
			},
		)
		store.Set(counterKey, idValue)
		store.Set(Key, value)
	}
	return nil
}

func migrateValue(oldVal v2types.Asset) (newVal types.Asset, oldKey []byte) {
	fmt.Println("migrateValue")
	asset := types.Asset{
		Id:        oldVal.Id,
		Name:      oldVal.Name,
		Denom:     oldVal.Denom,
		Decimal:   oldVal.Decimal,
		IbcStatus: oldVal.IbcStatus,
	}

	return asset, types.AssetKey(asset.Id)
}
