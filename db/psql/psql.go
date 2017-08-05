package psql

import (
	"errors"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	Db    *gorm.DB
	Error error
)

const dbEnvVar = "LUB_API_DATABASE_URL"

func Connect(envVar string) {
	u := os.Getenv(dbEnvVar)
	if u == "" {
		Error = errors.New("Missing " + dbEnvVar)
		return
	}

	Db, Error = gorm.Open("postgres", u)
}
