package main

import (
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"mmx.com/go-mmx/app"
	mmxtypes "mmx.com/go-mmx/types"
	"mmx.com/go-mmx/version"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	genaccscli "github.com/cosmos/cosmos-sdk/x/genaccounts/client/cli"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(mmxtypes.MmxBech32PrefixAccAddr, mmxtypes.MmxBech32PrefixAccPub)
	config.SetBech32PrefixForValidator(mmxtypes.MmxBech32PrefixValAddr, mmxtypes.MmxBech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(mmxtypes.MmxBech32PrefixConsAddr, mmxtypes.MmxBech32PrefixConsPub)
	config.SetCoinType(mmxtypes.MmxCoinType)
	config.SetFullFundraiserPath(mmxtypes.MmxFullFundraiserPath)
	config.Seal()

	ctx := server.NewDefaultContext()

	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "mmxd",
		Short:             "Mmx Chain Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	// CLI commands to initialize the chain
	rootCmd.AddCommand(genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.MigrateGenesisCmd(ctx, cdc))
	rootCmd.AddCommand(genutilcli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{},
		genaccounts.AppModuleBasic{}, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics))
	rootCmd.AddCommand(genaccscli.AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
	// rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))
	rootCmd.AddCommand(version.Cmd())

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "MMXD", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod, 0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewMmxApp(logger, db, traceStore, true, invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString("min-gas-prices")),
		baseapp.SetHaltHeight(uint64(viper.GetInt("halt-height"))),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		MmxApp := app.NewMmxApp(logger, db, traceStore, false, uint(1))
		err := MmxApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return MmxApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}
	MmxApp := app.NewMmxApp(logger, db, traceStore, true, uint(1))
	return MmxApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
