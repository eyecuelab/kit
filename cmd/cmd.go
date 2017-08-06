package cmd

import (
	"github.com/eyecuelab/kit/config"
	"github.com/eyecuelab/kit/log"
	"github.com/spf13/cobra"
)

var (
	Root     *cobra.Command
	commands []*cobra.Command
	verbose  bool
	cfgFile  string
	NoDb     bool
)

func init() {
	cobra.OnInitialize(initConfig)
}

func NewRoot(use string, short string, long string) *cobra.Command {
	Root = &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
	}

	Root.PersistentFlags().BoolVar(&verbose, "verbose", false, "more verbose error reporting")
	Root.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lub-api.yaml)")
	Root.PersistentFlags().BoolVar(&NoDb, "nodb", false, "allow DB-less execution")

	addCommands()
	return Root
}

func Add(command *cobra.Command) {
	commands = append(commands, command)
}

func addCommands() {
	for _, command := range commands {
		Root.AddCommand(command)
	}
}

func initConfig() {
	if err := config.Load("LUB_API", cfgFile); err != nil {
		log.FatalWrap(err, "Failed to load configuration")
	}
}
