package psql

import (
	"errors"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	DB      *gorm.DB
	DBError error
)

func init() {
	cobra.OnInitialize(ConnectDB)
}

func ConnectDB() {
	viper.SetDefault("database_scheme", "postgres")
	scheme := viper.GetString("database_scheme")
	url := viper.GetString("database_url")

	if len(url) == 0 {
		DBError = errors.New("Missing database_url")
	} else {
		if scheme != "postgres" {
			gorm.RegisterDialect(scheme, gorm.DialectsMap["postgres"])
		}
		DB, DBError = gorm.Open(scheme, url)
	}
}
