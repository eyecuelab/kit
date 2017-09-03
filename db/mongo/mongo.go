package mongo

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

var (
	MDb      *mgo.Database
	Error    error
	MongoUrl string
	Msession *mgo.Session
)

func init() {
	cobra.OnInitialize(connect)
}

func connect() {
	MongoUrl = viper.GetString("mongo_url")
	if len(MongoUrl) == 0 {
		Error = errors.New("Missing mongo_url")
	} else {
		if Msession, Error = mgo.Dial(MongoUrl); Error == nil {
			info, _ := mgo.ParseURL(MongoUrl)
			Msession.SetMode(mgo.Monotonic, true)
			MDb = Msession.DB(info.Database)
		}
	}
}
