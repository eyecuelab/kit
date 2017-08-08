package cmd

import (
	"github.com/eyecuelab/kit/log"
	"github.com/eyecuelab/kit/web"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//ApiCmd represents the api command
var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run the API server",
	Long:  ``,
	Run:   run,
}

var Port int

func init() {
	ApiCmd.PersistentFlags().IntVar(&Port, "port", 3000, "port to attach server")
	viper.BindPFlag("port", ApiCmd.PersistentFlags().Lookup("port"))
	viper.BindEnv("port")
}

func run(cmd *cobra.Command, args []string) {
	if viper.GetInt("port") > 0 {
		Port = viper.GetInt("port")
	}
	log.Infof("Serving API on port %d...", Port)
	web.Start(Port)
}
