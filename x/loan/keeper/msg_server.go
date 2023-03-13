package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/loan/types"
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

func (m msgServer) MsgCreate(goCtx context.Context, request *types.MsgCreateRequest) (*types.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.Keeper.Request(ctx, request.AppId, request.PairId, request.GetFrom(), request.AmountIn, request.AmountOut, request.Fee, request.DurationDays)
	if err != nil {
		return nil, err
	}
	return &types.MsgCreateResponse{}, nil
}

func (m msgServer) MsgApprove(goCtx context.Context, request *types.MsgApproveRequest) (*types.MsgApproveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.Keeper.Approve(ctx, request.LoanId, request.GetFrom())
	if err != nil {
		return nil, err
	}
	return &types.MsgApproveResponse{}, nil
}

func (m msgServer) MsgRepay(goCtx context.Context, request *types.MsgRepayRequest) (*types.MsgRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.Keeper.Repay(ctx, request.LoanId)
	if err != nil {
		return nil, err
	}
	return &types.MsgRepayResponse{}, nil
}

func (m msgServer) MsgCancel(goCtx context.Context, request *types.MsgCancelRequest) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.Keeper.Cancel(ctx, request.LoanId)
	if err != nil {
		return nil, err
	}
	return &types.MsgCancelResponse{}, nil
}
