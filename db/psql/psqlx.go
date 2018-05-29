package psql

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/eyecuelab/kit/flect"
	"github.com/eyecuelab/kit/islice"
	"github.com/eyecuelab/kit/stringslice"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	//register postgres dialect
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/lib/pq"
)

var (
	DBx      *IQSqlx
	DBxError error
)

func init() {
	cobra.OnInitialize(ConnectDBx)
}

type (
	IQSqlx struct {
		X *sqlx.DB
	}

	IQModel interface {
		TableName() string
	}

	Mappable interface {
		IdColumn() string
	}
)

func ConnectDBx() {
	viper.SetDefault("database_scheme", "postgres")
	scheme := viper.GetString("database_scheme")
	url := viper.GetString("database_url")

	if len(url) == 0 {
		DBxError = errors.New("Missing database_url")
		return
	}
	var dbx *sqlx.DB
	dbx, DBxError = sqlx.Connect(scheme, url)
	DBx = &IQSqlx{dbx}
}


//Takes a map[string]struct, populates the stuct, and sets the map keys to the column specified by the mappable interface
func (db IQSqlx) MapById(mappable Mappable, query string, params ...interface{}) error {
	if flect.NotA(mappable, reflect.Map) {
		return fmt.Errorf("MapById: mappable must be a map, %s is a %T", reflect.TypeOf(mappable).Name(), mappable)
	}

	rows, err := db.X.Queryx(query, params...)
	if err != nil {
		return err
	}

	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	idColIndex := stringslice.IndexOf(cols, mappable.IdColumn())

	valuePtrs := islice.StringPtrs(len(cols))

	for rows.Next() {
		if err := rows.Scan(valuePtrs...); err != nil {
			return err
		}

		mapDeref := reflect.TypeOf(mappable).Elem()

		if mapDeref.Kind() == reflect.Map {
			structInterface := reflect.New(mapDeref.Elem()).Interface()
			if err := rows.StructScan(structInterface); err != nil {
				return err
			}

			key := reflect.ValueOf(valuePtrs[idColIndex]).Elem()
			value := reflect.ValueOf(structInterface).Elem()

			reflect.ValueOf(mappable).Elem().SetMapIndex(key, value)
		}
	}
	return nil
}
