package psql

import (
	"github.com/jinzhu/gorm"

	//register postgres dialect

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	testDBName = "test"
	postgres   = "postgres"
	url        = `postgresql://localhost/test?sslmode=disable`
)

func TestDB() (*gorm.DB, error) {
	return gorm.Open(postgres, url)
}
