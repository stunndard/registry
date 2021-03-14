package registry

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stunndard/registry/x/registry/keeper"
	"github.com/stunndard/registry/x/registry/types"
	//"github.com/tendermint/tendermint/crypto"
)

func handleMsgBuyName(ctx sdk.Context, k keeper.Keeper, msg *types.MsgBuyName) (*sdk.Result, error) {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid buyer address")
	}

	// if this name is registered
	if k.HasName(ctx, msg.Name) {
		// fetch name
		name := k.GetName(ctx, msg.Name)

		// if creator is already owner of this name
		if msg.Creator == name.Owner {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Already owner")
		}

		// if this name is NOT for sale by the current owner
		if !name.Onsale {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name not on sale")
		}

		// get the current name price
		price := k.GetPrice(ctx, msg.Name)

		// if bid is less than or equal to name price
		if msg.Bid.IsAllLTE(price) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough")
		}

		to, err := sdk.AccAddressFromBech32(name.Owner)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid owner address")
		}

		// send bid amount of tokens from msg creator to the current owner
		if err := k.CoinKeeper.SendCoins(ctx, from, to, msg.Bid); err != nil {
			return nil, err
		}

		// set name owner
		name.Owner = msg.Creator

		// set name price
		name.Price = msg.Bid

		// set name onsale
		name.Onsale = msg.Onsale

		k.SetName(ctx, name)

	} else {
		// this name is not registered
		// TODO: amount can be calculated dynamically from the total number of names registered for the N latest blocks
		amount, _ := sdk.ParseCoinsNormalized("10token")
		if err := k.CoinKeeper.SubtractCoins(ctx, from, amount); err != nil {
			return nil, err
		}

		k.CreateName(ctx, *msg)
	}

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
