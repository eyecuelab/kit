package psql

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	//register postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Scope func(db *gorm.DB) *gorm.DB
type JsonB map[string]interface{}

const nullDataValue = "null"

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
		DB, DBError = gorm.Open(scheme, url)
		// DB.LogMode(true)
	}
}

func (j JsonB) Value() (driver.Value, error) {
	return json.Marshal(j)
}
func (j *JsonB) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	if sourceIsNull(source) {
		return nil
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*j, ok = i.(map[string]interface{})
	if !ok {
		return fmt.Errorf("type assertion .(map[string]interface{}) failed. got %s", source)
	}

	return nil
}

func UndeleteOrCreate(model interface{}, query string, args ...interface{}) error {
	tx := DB.Begin()
	tx = tx.Unscoped().Model(model).Where(query, args...).UpdateColumn("deleted_at", nil)

	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	if tx.RowsAffected > 1 {
		tx.Rollback()
		return fmt.Errorf("UndeleteOrCreate can only undelete one record, you're trying to undelete %v", tx.RowsAffected)
	}

	if tx.RowsAffected < 1 {
		if err := tx.Create(model).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
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

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	if tx.RowsAffected < 1 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func Transact(db *gorm.DB, txFunc func(*gorm.DB) error) (err error) {
	tx := db.Begin()
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}

func DbWithError(e error) *gorm.DB {
	db := DB.New()
	db.Error = e
	return db
}

func sourceIsNull(b []byte) bool {
	return string(b) == nullDataValue
}
