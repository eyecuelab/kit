package mongo

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

var (
	MDb      *mgo.Session
	Error    error
	MongoUrl string
)

func init() {
	cobra.OnInitialize(connect)
}

func connect() {
	MongoUrl = viper.GetString("mongo_url")
	if len(MongoUrl) == 0 {
		Error = errors.New("Missing mongo_url")
	} else {
		if MDb, Error = mgo.Dial(MongoUrl); Error == nil {
			MDb.SetMode(mgo.Monotonic, true)
		}
	}
}
