package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateRequest{}, "comdex/loan/MsgCreateRequest", nil)
	cdc.RegisterConcrete(&MsgApproveRequest{}, "comdex/loan/MsgApproveRequest", nil)
	cdc.RegisterConcrete(&MsgRepayRequest{}, "comdex/loan/MsgRepayRequest", nil)
	cdc.RegisterConcrete(&MsgCancelRequest{}, "comdex/loan/MsgCancelRequest", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateRequest{},
		&MsgApproveRequest{},
		&MsgRepayRequest{},
		&MsgCancelRequest{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
