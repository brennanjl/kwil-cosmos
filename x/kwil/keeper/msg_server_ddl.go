package keeper

import (
	"context"

	"kwil/x/kwil/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DDL(goCtx context.Context, msg *types.MsgDDL) (*types.MsgDDLResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// First check if the specified DB exists
	_, isFound := k.GetDatabases(ctx, msg.Dbid)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Specified DBid does not exist")
	}

	if !k.IsDbOwner(ctx, msg.Dbid, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Caller is not the owner of this DB")
	}

	return &types.MsgDDLResponse{}, nil
}
