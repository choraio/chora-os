package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateSlog{}, "ecodex/CreateSlog", nil)
	cdc.RegisterConcrete(&MsgUpdateSlog{}, "ecodex/UpdateSlog", nil)
	cdc.RegisterConcrete(&MsgDeleteSlog{}, "ecodex/DeleteSlog", nil)

	cdc.RegisterConcrete(&MsgCancelBuyOrder{}, "ecodex/CancelBuyOrder", nil)

	cdc.RegisterConcrete(&MsgCancelSellOrder{}, "ecodex/CancelSellOrder", nil)

	cdc.RegisterConcrete(&MsgSendBuyOrder{}, "ecodex/SendBuyOrder", nil)

	cdc.RegisterConcrete(&MsgSendSellOrder{}, "ecodex/SendSellOrder", nil)

	cdc.RegisterConcrete(&MsgSendCreatePair{}, "ecodex/SendCreatePair", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSlog{},
		&MsgUpdateSlog{},
		&MsgDeleteSlog{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelBuyOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelSellOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendBuyOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendSellOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendCreatePair{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
