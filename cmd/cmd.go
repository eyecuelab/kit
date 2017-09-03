package cmd

import (
	"github.com/eyecuelab/kit/assets"
	"github.com/eyecuelab/kit/cmd/api"
	"github.com/eyecuelab/kit/cmd/migrate"
	"github.com/eyecuelab/kit/config"
	"github.com/spf13/cobra"
)

var (
	Root              *cobra.Command
	childCommands     []*cobra.Command
	availableCommands = make(map[string]*cobra.Command)
	verbose           bool
	cfgFile           string
	NoDb              bool
)

func Add(command *cobra.Command) {
	childCommands = append(childCommands, command)
}

func Use(cmds ...string) {
	for _, key := range cmds {
		Add(availableCommands[key])
	}
}

func Init(appName string, rootCmd *cobra.Command, assetGet assets.AssetGet, assetDir assets.AssetDir) error {
	assets.Manager = &assets.AssetManager{assetGet, assetDir}
	addRoot(rootCmd)

	return config.Load(appName, cfgFile)
}

func addRoot(cmd *cobra.Command) {
	Root = cmd

	Root.PersistentFlags().BoolVar(&verbose, "verbose", false, "more verbose error reporting")
	Root.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/config.yaml)")
	Root.PersistentFlags().BoolVar(&NoDb, "nodb", false, "allow DB-less execution")

	Root.AddCommand(childCommands...)
}

func init() {
	availableCommands["api"] = api.ApiCmd
	availableCommands["migrate"] = migrate.MigrateCmd
}
