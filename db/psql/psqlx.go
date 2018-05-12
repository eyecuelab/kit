package psql

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	//register postgres dialect
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/lib/pq"
)

const (
	selectAll  = "select * from %s"
	selectFind = "select * from %s where id = ?"
)

var (
	DBx      *IQSqlx
	DBxError error
)

func init() {
	cobra.OnInitialize(ConnectDBx)
}

type IQSqlx struct {
	X *sqlx.DB
}

type IQModel interface {
	TableName() string
}

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

func (iqx IQSqlx) All(m interface{}, ref IQModel) error {
	clause := fmt.Sprintf(selectAll, ref.TableName())
	return iqx.X.Select(m, clause)
}

func (iqx IQSqlx) Find(m IQModel, id string) error {
	clause := fmt.Sprintf(selectFind, m.TableName())
	return iqx.X.Get(m, clause, id)
}
