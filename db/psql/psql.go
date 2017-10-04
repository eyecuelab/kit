package psql

import (
	"errors"

	"database/sql/driver"
	"encoding/json"

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
		// TODO: seems like a bug, breaks the build for me
		if scheme != "postgres" {
			gorm.RegisterDialect(scheme, gorm.DialectsMap["postgres"])
		}
		DB, DBError = gorm.Open(scheme, url)
		DB.LogMode(true)
	}
}

type JsonB map[string]interface{}

func (j JsonB) Value() (driver.Value, error) {
	return json.Marshal(j)
}
func (j *JsonB) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*j, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}

func DeleteModelWithAssociations(value interface{}, associations ...string) error {
	tx := DB.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	for _, a := range associations {
		if err := tx.Model(value).Association(a).Clear().Error; err != nil {
			return err
		}
	}

	tx = tx.Delete(value)

	if tx.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
