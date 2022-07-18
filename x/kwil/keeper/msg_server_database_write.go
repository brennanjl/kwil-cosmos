package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"kwil/x/kwil/types"
)

func (k msgServer) DatabaseWrite(goCtx context.Context, msg *types.MsgDatabaseWrite) (*types.MsgDatabaseWriteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDatabaseWriteResponse{}, nil
}
