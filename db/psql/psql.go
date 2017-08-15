package psql

import (
	"errors"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/eyecuelab/kit/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Db         *gorm.DB
	DbUrl      string
	Scheme     string
	Error      error
	references []*gorm.DB
)

func init() {
	cobra.OnInitialize(connect)
	viper.SetDefault("database_scheme", "postgres")
}

func connect() {
	Scheme = viper.GetString("database_scheme")
	DbUrl = viper.GetString("database_url")

	if DbUrl == "" {
		Error = errors.New("Missing database url")
	} else {
		if Scheme != "postgres" {
			log.Infof("Registering dialect: %s", Scheme)
			gorm.RegisterDialect(Scheme, gorm.DialectsMap["postgres"])
		}
		Db, Error = gorm.Open(Scheme, DbUrl)
	}
}
