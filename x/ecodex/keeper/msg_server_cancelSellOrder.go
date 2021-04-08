package keeper

import (
	"context"
	"errors"

	"github.com/choraio/chora/x/ecodex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CancelSellOrder(goCtx context.Context, msg *types.MsgCancelSellOrder) (*types.MsgCancelSellOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve the book
	pairIndex := types.OrderBookIndex(msg.Port, msg.Channel, msg.AmountDenom, msg.PriceDenom)
	book, found := k.GetSellOrderBook(ctx, pairIndex)
	if !found {
		return &types.MsgCancelSellOrderResponse{}, errors.New("the pair doesn't exist")
	}

	// Check order creator
	order, err := book.GetOrderFromID(msg.OrderID)
	if err != nil {
		return &types.MsgCancelSellOrderResponse{}, err
	}
	if order.Creator != msg.Creator {
		return &types.MsgCancelSellOrderResponse{}, errors.New("canceller must be creator")
	}

	// Remove order
	newBook, err := book.RemoveOrderFromID(msg.OrderID)
	if err != nil {
		return &types.MsgCancelSellOrderResponse{}, err
	}
	book = newBook.(types.SellOrderBook)
	k.SetSellOrderBook(ctx, book)

    // Refund seller with remaining amount
	seller, err := sdk.AccAddressFromBech32(order.Creator)
	if err != nil {
		return &types.MsgCancelSellOrderResponse{}, err
	}
	if err := k.SafeMint(
		ctx,
		msg.Port,
		msg.Channel,
		seller,
		msg.AmountDenom,
		order.Amount,
	); err != nil {
		return &types.MsgCancelSellOrderResponse{}, err
	}

	return &types.MsgCancelSellOrderResponse{}, nil
}
