package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = QueryServer{}

type QueryServer struct {
	Keeper
}

func NewQueryServer(k Keeper) types.QueryServer {
	return &QueryServer{
		Keeper: k,
	}
}

func (q QueryServer) QueryLoan(c context.Context, req *types.QueryLoanRequest) (*types.QueryLoanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetLoan(ctx, req.Id)
	if !found {
		return &types.QueryLoanResponse{}, nil
	}

	return &types.QueryLoanResponse{
		Loan: item,
	}, nil
}
