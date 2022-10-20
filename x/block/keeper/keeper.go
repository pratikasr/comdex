package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	v2types "github.com/comdex-official/comdex/x/block/migrations/v2/types"
	"github.com/comdex-official/comdex/x/block/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	protobuftypes "github.com/gogo/protobuf/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
func (k Keeper) AddAsset(ctx sdk.Context, msg types.MsgAddAsset) error {
	id := k.NewGetAssetID(ctx)
	asset := types.Asset{
		Id:      id + 1,
		Name:    msg.Name,
		Denom:   msg.Denom,
		Decimal: msg.Decimal,
		//Price:     msg.Price,
		//AppId:     msg.AppId,
		IbcStatus: msg.IbcStatus,
	}
	k.NewSetAssetID(ctx, asset.Id)
	k.NewSetAsset(ctx, asset)
	return nil
}

func (k Keeper) GetAssetID(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.AssetIDKey
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

func (k Keeper) GetAsset(ctx sdk.Context, id uint64) (asset v2types.Asset, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.AssetKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return asset, false
	}

	k.cdc.MustUnmarshal(value, &asset)
	return asset, true
}

func (k Keeper) NewSetAssetID(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.NewAssetIDKey
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)

	store.Set(key, value)
}

func (k Keeper) NewGetAssetID(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.NewAssetIDKey
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k Keeper) NewSetAsset(ctx sdk.Context, asset types.Asset) {
	var (
		store = k.Store(ctx)
		key   = types.AssetKey(asset.Id)
		value = k.cdc.MustMarshal(&asset)
	)

	store.Set(key, value)
}

func (k Keeper) NewGetAsset(ctx sdk.Context, id uint64) (asset types.Asset, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.AssetKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return asset, false
	}

	k.cdc.MustUnmarshal(value, &asset)
	return asset, true
}

func (k Keeper) MigrateValuesAsset(ctx sdk.Context) error {
	assetID := k.GetAssetID(ctx)
	fmt.Println("assetID", assetID)
	for i := 1; i <= int(assetID); i++ {
		asset, _ := k.GetAsset(ctx, uint64(i))
		fmt.Println("oldAsset", asset)
		msg := types.MsgAddAsset{
			Name:      asset.Name,
			Denom:     asset.Denom,
			Decimal:   asset.Decimal,
			IbcStatus: asset.IbcStatus,
		}
		err := k.AddAsset(ctx, msg)
		fmt.Println("err", err)
		if err != nil {
			return err
		}
	}
	return nil
}
