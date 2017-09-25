package api

import (
	"github.com/eyecuelab/kit/cmd/migrate"
	"github.com/eyecuelab/kit/log"
	"github.com/eyecuelab/kit/web/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var apiArgs = []string{"port", "secret"}

//ApiCmd represents the api command
var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run the API server",
	Long:  ``,
	Run:   run,
}

var (
	Port      int
	Domain    string
	checkMigs bool
)

func init() {
	ApiCmd.PersistentFlags().IntVar(&Port, "port", 3000, "port to attach server")
	ApiCmd.PersistentFlags().String("secret", "", "secret key used for token hashing")
	ApiCmd.PersistentFlags().BoolVar(&checkMigs, "check-migrations", true, "check pending migrations before starting server")
	ApiCmd.PersistentFlags().StringVar(&Domain, "domain", "", "the url to reach this api eg, https://foo.ngrok.io")

	for _, a := range apiArgs {
		viper.BindPFlag(a, ApiCmd.PersistentFlags().Lookup(a))
		viper.BindEnv(a)
	}
}

func run(cmd *cobra.Command, args []string) {
	if checkMigs {
		checkMigrations()
	}

	if viper.GetInt("port") > 0 {
		Port = viper.GetInt("port")
	}
	log.Infof("Serving API on port %d...", Port)

	server.Start(Port, Domain)
}

func checkMigrations() {
	if c, err := migrate.PendingMigrationCount(); err != nil {
		log.Fatal(err)
	} else if c > 0 {
		log.Fatalf("%d Pending Migration(s), run migrations or use check-migrations=0", c)
	}
}
