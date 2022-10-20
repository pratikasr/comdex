package block

import (
	"fmt"
	utils "github.com/comdex-official/comdex/types"
	"github.com/comdex-official/comdex/x/block/keeper"
	"github.com/comdex-official/comdex/x/block/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, ctx.BlockTime(), telemetry.MetricKeyBeginBlocker)

	_ = utils.ApplyFuncIfNoError(ctx, func(ctx sdk.Context) error {
		fmt.Println(ctx.BlockHeight())
		if ctx.BlockHeight() == 35 {
			fmt.Println("ctx.BlockHeight() == 35")
			err := k.MigrateValuesAsset(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
