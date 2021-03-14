package registry

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stunndard/registry/x/registry/keeper"
	"github.com/stunndard/registry/x/registry/types"
)

func handleMsgUpdateName(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateName) (*sdk.Result, error) {

	// Checks that the element exists
	if !k.HasName(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Name %s doesn't exist", msg.Name))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetNameOwner(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	name := types.Name{
		Creator: k.GetNameCreator(ctx, msg.Name),
		Id:      k.GetName(ctx, msg.Name).Id,
		Name:    k.GetName(ctx, msg.Name).Name,
		Owner:   msg.Creator,
		Price:   msg.Price,
		Onsale:  msg.Onsale,
	}

	k.SetName(ctx, name)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateName,
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyNewOwner, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Price.String()),
			sdk.NewAttribute(types.AttributeKeyOnSale, strconv.FormatBool(msg.Onsale)),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
