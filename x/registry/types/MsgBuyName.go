package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBuyName{}

func NewMsgBuyName(creator string, name string, bid sdk.Coins, onsale bool) *MsgBuyName {
	return &MsgBuyName{
		Creator: creator,
		Name:    name,
		Bid:     bid,
		Onsale:  onsale,
	}
}

func (msg *MsgBuyName) Route() string {
	return RouterKey
}

func (msg *MsgBuyName) Type() string {
	return "BuyName"
}

func (msg *MsgBuyName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
