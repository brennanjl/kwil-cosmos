package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDefineQuery = "define_query"

var _ sdk.Msg = &MsgDefineQuery{}

func NewMsgDefineQuery(creator string, parQuer string, publicity bool) *MsgDefineQuery {
	return &MsgDefineQuery{
		Creator:   creator,
		ParQuer:   parQuer,
		Publicity: publicity,
	}
}

func (msg *MsgDefineQuery) Route() string {
	return RouterKey
}

func (msg *MsgDefineQuery) Type() string {
	return TypeMsgDefineQuery
}

func (msg *MsgDefineQuery) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDefineQuery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDefineQuery) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
