package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"kwil/x/kwil/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDatabase(goCtx context.Context, msg *types.MsgCreateDatabase) (*types.MsgCreateDatabaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// First, must combine the sender and block height.
	// TODO: Right now the uniqueness is determined by a seed that is provided by the user.
	// This is not an ideal system, and is only used for prototyping right now.
	dbIDBytes := sha256.Sum256([]byte(msg.Creator + msg.Seed))

	// Type cast to string
	dbName := hex.EncodeToString(dbIDBytes[:])

	// Check if db ID already exists
	_, isFound := k.GetDatabases(ctx, dbName)

	// Return an error if it exists
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Database for that user with that seed already exists, please use a different seed")
	}

	// Create the new DB
	newDB := types.Databases{
		Index: dbName,
		Dbid:  dbName,
		Owner: msg.Creator,
	}

	// Store the DB
	k.SetDatabases(ctx, newDB)

	return &types.MsgCreateDatabaseResponse{Id: dbName}, nil
}
