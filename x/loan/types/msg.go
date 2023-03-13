package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgCreate(addr string, appID, pairID uint64, amountIn, amountOut, fee sdk.Int, durationDays int64) *MsgCreateRequest {
	return &MsgCreateRequest{
		From:         addr,
		AppId:        appID,
		PairId:       pairID,
		AmountIn:     amountIn,
		AmountOut:    amountOut,
		Fee:          fee,
		DurationDays: durationDays,
	}
}

func (msg MsgCreateRequest) Route() string { return ModuleName }
func (msg MsgCreateRequest) Type() string  { return TypeCreateRequest }

func (msg *MsgCreateRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetFrom())
	if err != nil {
		return err
	}
	if msg.AppId == 0 {
		return fmt.Errorf("app id should not be 0: %d ", msg.AppId)
	}
	if msg.AmountIn.IsNegative() || msg.AmountIn.IsZero() {
		return fmt.Errorf("invalid coin amount: %s < 0", msg.AmountIn)
	}
	if msg.AmountOut.IsNegative() || msg.AmountOut.IsZero() {
		return fmt.Errorf("invalid coin amount: %s < 0", msg.AmountOut)
	}
	if msg.PairId == 0 {
		return fmt.Errorf("pair id should not be 0: %d ", msg.PairId)
	}

	return nil
}

func (msg *MsgCreateRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.GetFrom())
	return []sdk.AccAddress{addr}
}

// GetSignBytes get the bytes for the message signer to sign on.
func (msg *MsgCreateRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func NewMsgApprove(addr string, loanID uint64) *MsgApproveRequest {
	return &MsgApproveRequest{
		From:   addr,
		LoanId: loanID,
	}
}

func (msg MsgApproveRequest) Route() string { return ModuleName }
func (msg MsgApproveRequest) Type() string  { return TypeCreateRequest }

func (msg *MsgApproveRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetFrom())
	if err != nil {
		return err
	}

	return nil
}

func (msg *MsgApproveRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.GetFrom())
	return []sdk.AccAddress{addr}
}

// GetSignBytes get the bytes for the message signer to sign on.
func (msg *MsgApproveRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func NewMsgRepay(addr string, loanID uint64) *MsgRepayRequest {
	return &MsgRepayRequest{
		From:   addr,
		LoanId: loanID,
	}
}

func (msg MsgRepayRequest) Route() string { return ModuleName }
func (msg MsgRepayRequest) Type() string  { return TypeCreateRequest }

func (msg *MsgRepayRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetFrom())
	if err != nil {
		return err
	}

	return nil
}

func (msg *MsgRepayRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.GetFrom())
	return []sdk.AccAddress{addr}
}

// GetSignBytes get the bytes for the message signer to sign on.
func (msg *MsgRepayRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func NewMsgCancel(addr string, loanID uint64) *MsgCancelRequest {
	return &MsgCancelRequest{
		From:   addr,
		LoanId: loanID,
	}
}

func (msg MsgCancelRequest) Route() string { return ModuleName }
func (msg MsgCancelRequest) Type() string  { return TypeCreateRequest }

func (msg *MsgCancelRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetFrom())
	if err != nil {
		return err
	}

	return nil
}

func (msg *MsgCancelRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.GetFrom())
	return []sdk.AccAddress{addr}
}

// GetSignBytes get the bytes for the message signer to sign on.
func (msg *MsgCancelRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
