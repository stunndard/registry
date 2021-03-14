package registry

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stunndard/registry/x/registry/keeper"
	"github.com/stunndard/registry/x/registry/types"
)

func handleMsgDeleteName(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteName) (*sdk.Result, error) {
	if !k.HasName(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Name %s doesn't exist", msg.Name))
	}
	if msg.Creator != k.GetNameOwner(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.DeleteName(ctx, msg.Name)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
