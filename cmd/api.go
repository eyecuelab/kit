package cmd

import (
	"github.com/eyecuelab/kit/log"
	"github.com/eyecuelab/kit/web"

	"github.com/spf13/cobra"
)

//ApiCmd represents the api command
var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run the API server",
	Long:  ``,
	Run:   run,
}

var (
	port int
)

func init() {
	ApiCmd.PersistentFlags().IntVar(&port, "port", 3000, "port to attach server")
}

func run(cmd *cobra.Command, args []string) {
	log.Infof("Serving API on port %d...", port)
	web.Start(port)
}
