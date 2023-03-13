package keeper

import (
	"github.com/comdex-official/comdex/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuftypes "github.com/gogo/protobuf/types"
)

func (k Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

func (k Keeper) SetLoanID(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LoanIDKey
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)

	store.Set(key, value)
}

func (k Keeper) GetLoanID(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.LoanIDKey
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k Keeper) SetLoan(ctx sdk.Context, loan types.Loan) {
	var (
		store = k.Store(ctx)
		key   = types.LoanKey(loan.Id)
		value = k.cdc.MustMarshal(&loan)
	)

	store.Set(key, value)
}

func (k Keeper) GetLoan(ctx sdk.Context, id uint64) (loan types.Loan, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.LoanKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return loan, false
	}

	k.cdc.MustUnmarshal(value, &loan)
	return loan, true
}

func (k Keeper) DeleteLoan(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LoanKey(id)
	)

	store.Delete(key)
}
