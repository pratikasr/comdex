package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/loan module sentinel errors
var (
	ErrWrongLoanState = sdkerrors.Register(ModuleName, 1100, "Wrong loan state")
)
