package psql

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/jmoiron/sqlx"

	//register postgres dialect
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/lib/pq"
)

var (
	DBx      *sqlx.DB
	DBxError error
)

func init() {
	cobra.OnInitialize(ConnectDBx)
}

func ConnectDBx() {
	viper.SetDefault("database_scheme", "postgres")
	scheme := viper.GetString("database_scheme")
	url := viper.GetString("database_url")

	if len(url) == 0 {
		DBxError = errors.New("Missing database_url")
		return
	}
	DBx, DBxError = sqlx.Connect(scheme, url)
	// DB.LogMode(true)
}