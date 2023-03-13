package keeper

import (
	"fmt"
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/comdex-official/comdex/x/loan/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
		bankKeeper types.BankKeeper
		Asset      types.AssetKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	asset types.AssetKeeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bankKeeper,
		Asset:      asset,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Request(ctx sdk.Context, appID, pairID uint64, owner string, amtIn, amtOut, fee sdk.Int, duration int64) error {
	loanID := k.GetLoanID(ctx)
	loan := types.Loan{
		Id:           loanID + 1,
		AppId:        appID,
		PairID:       pairID,
		Owner:        owner,
		AmountIn:     amtIn,
		AmountOut:    amtOut,
		Fee:          fee,
		CreatedAt:    ctx.BlockTime(),
		DurationDays: duration,
		Lender:       "",
		State:        "requested",
	}
	addr, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return err
	}
	pair, found := k.Asset.GetPair(ctx, pairID)
	if !found {
		return assettypes.ErrorPairDoesNotExist
	}
	assetIn, found := k.Asset.GetAsset(ctx, pair.AssetIn)
	if !found {
		return assettypes.ErrorAssetDoesNotExist
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(sdk.NewCoin(assetIn.Denom, amtIn.Add(fee))))
	if err != nil {
		return err
	}
	k.SetLoanID(ctx, loan.Id)
	k.SetLoan(ctx, loan)
	return nil
}

func (k Keeper) Approve(ctx sdk.Context, loanID uint64, lender string) error {
	loan, found := k.GetLoan(ctx, loanID)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist for loanID", loanID)
	}
	if loan.State != "requested" {
		return sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}

	pair, found := k.Asset.GetPair(ctx, loan.PairID)
	if !found {
		return assettypes.ErrorPairDoesNotExist
	}

	assetOut, found := k.Asset.GetAsset(ctx, pair.AssetOut)
	if !found {
		return assettypes.ErrorAssetDoesNotExist
	}

	lenderAddr, err := sdk.AccAddressFromBech32(lender)
	if err != nil {
		return err
	}

	borrowerAddr, err := sdk.AccAddressFromBech32(loan.Owner)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, lenderAddr, types.ModuleName, sdk.NewCoins(sdk.NewCoin(assetOut.Denom, loan.AmountOut)))
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrowerAddr, sdk.NewCoins(sdk.NewCoin(assetOut.Denom, loan.AmountOut)))
	if err != nil {
		return err
	}
	loan.Lender = lender
	loan.State = "approved"
	return nil
}

func (k Keeper) Repay(ctx sdk.Context, loanID uint64) error {
	loan, found := k.GetLoan(ctx, loanID)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist for loanID", loanID)
	}
	if loan.State != "approved" {
		return sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}

	pair, found := k.Asset.GetPair(ctx, loan.PairID)
	if !found {
		return assettypes.ErrorPairDoesNotExist
	}

	assetIn, found := k.Asset.GetAsset(ctx, pair.AssetIn)
	if !found {
		return assettypes.ErrorAssetDoesNotExist
	}

	assetOut, found := k.Asset.GetAsset(ctx, pair.AssetOut)
	if !found {
		return assettypes.ErrorAssetDoesNotExist
	}

	lenderAddr, err := sdk.AccAddressFromBech32(loan.Lender)
	if err != nil {
		return err
	}

	borrowerAddr, err := sdk.AccAddressFromBech32(loan.Owner)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrowerAddr, types.ModuleName, sdk.NewCoins(sdk.NewCoin(assetOut.Denom, loan.AmountOut)))
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lenderAddr, sdk.NewCoins(sdk.NewCoin(assetOut.Denom, loan.AmountOut), sdk.NewCoin(assetIn.Denom, loan.Fee)))
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrowerAddr, sdk.NewCoins(sdk.NewCoin(assetIn.Denom, loan.AmountIn)))
	if err != nil {
		return err
	}

	k.DeleteLoan(ctx, loanID)

	return nil
}

func (k Keeper) Cancel(ctx sdk.Context, loanID uint64) error {
	loan, found := k.GetLoan(ctx, loanID)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist for loanID", loanID)
	}
	if loan.State != "requested" {
		return sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}

	borrowerAddr, err := sdk.AccAddressFromBech32(loan.Owner)
	if err != nil {
		return err
	}

	pair, found := k.Asset.GetPair(ctx, loan.PairID)
	if !found {
		return assettypes.ErrorPairDoesNotExist
	}

	assetIn, found := k.Asset.GetAsset(ctx, pair.AssetIn)
	if !found {
		return assettypes.ErrorAssetDoesNotExist
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrowerAddr, sdk.NewCoins(sdk.NewCoin(assetIn.Denom, loan.AmountIn.Add(loan.Fee))))
	if err != nil {
		return err
	}
	k.DeleteLoan(ctx, loanID)

	return nil
}
