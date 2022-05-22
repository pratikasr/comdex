package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) Whitelist(goCtx context.Context, msg *types.WhitelistAsset) (*types.MsgWhitelistAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.WhitelistAsset(ctx, msg.AppMappingId, msg.AssetId); err != nil {
		return nil, err
	}
	return &types.MsgWhitelistAssetResponse{}, nil
}

func (m msgServer) RemoveWhitelist(goCtx context.Context, msg *types.RemoveWhitelistAsset) (*types.MsgRemoveWhitelistAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.RemoveWhitelistAsset(ctx, msg.AppMappingId, msg.AssetId); err != nil {
		return nil, err
	}
	return &types.MsgRemoveWhitelistAssetResponse{}, nil
}

func (m msgServer) WhitelistAppVault(goCtx context.Context, msg *types.WhitelistAppIdVault) (*types.MsgWhitelistAppIdVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.WhitelistAppIdVault(ctx, msg.AppMappingId); err != nil {
		return nil, err
	}
	return &types.MsgWhitelistAppIdVaultResponse{}, nil
}

func (m msgServer) RemoveWhitelistAppVault(goCtx context.Context, msg *types.RemoveWhitelistAppIdVault) (*types.MsgRemoveWhitelistAppIdVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.RemoveWhitelistAppIdVault(ctx, msg.AppMappingId); err != nil {
		return nil, err
	}
	return &types.MsgRemoveWhitelistAppIdVaultResponse{}, nil
}
