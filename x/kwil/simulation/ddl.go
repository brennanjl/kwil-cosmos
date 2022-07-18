package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"kwil/x/kwil/keeper"
	"kwil/x/kwil/types"
)

func SimulateMsgDDL(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDDL{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the DDL simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "DDL simulation not implemented"), nil, nil
	}
}
