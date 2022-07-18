package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"kwil/x/kwil/types"
)

func (k msgServer) DefineQuery(goCtx context.Context, msg *types.MsgDefineQuery) (*types.MsgDefineQueryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDefineQueryResponse{}, nil
}
