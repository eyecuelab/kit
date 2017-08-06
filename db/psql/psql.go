package psql

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Db         *gorm.DB
	Error      error
	references []*gorm.DB
)

func init() {
	cobra.OnInitialize(connect)
}

func connect() {
	dbUrl := viper.GetString("database_url")
	if dbUrl == "" {
		Error = errors.New("Missing database url")
	} else {
		Db, Error = gorm.Open("postgres", dbUrl)
	}
}
