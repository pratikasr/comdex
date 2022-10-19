package simulation

import (
	"math/rand"

	"github.com/comdex-official/comdex/x/block/keeper"
	"github.com/comdex-official/comdex/x/block/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddAsset(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddAsset{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddAsset simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddAsset simulation not implemented"), nil, nil
	}
}
