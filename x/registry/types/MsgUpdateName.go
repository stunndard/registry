package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateName{}

func NewMsgUpdateName(creator string, name string, owner string, price sdk.Coins, onsale bool) *MsgUpdateName {
	return &MsgUpdateName{
		Creator: creator,
		Name:    name,
		Owner:   owner,
		Price:   price,
		Onsale:  onsale,
	}
}

func (msg *MsgUpdateName) Route() string {
	return RouterKey
}

func (msg *MsgUpdateName) Type() string {
	return "UpdateName"
}

func (msg *MsgUpdateName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
