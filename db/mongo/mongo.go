package mongo

import (
	"errors"

	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

var (
	MDb      *mgo.Database
	Error    error
	MongoUrl string
	Msession *mgo.Session
)

func init() {
	cobra.OnInitialize(connectG)
}

func connectG() {
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

func Connect() (*mgo.Database, error){
	url := viper.GetString("mongo_url")
	if len(url) == 0 {
		return nil, errors.New("Missing mongo_url")
	}

	mSession, err := mgo.Dial(MongoUrl)
	if err != nil {
		return nil, err
	}

	info, _ := mgo.ParseURL(url)
	mSession.SetMode(mgo.Monotonic, true)

	return mSession.DB(info.Database), nil
}

//InCollection returns whether document(s) matching the query the specified collection exist
func InCollection(collection *mgo.Collection, selector interface{}) bool {
	n, _ := collection.Find(selector).Count()
	return n > 0
}

//UniqueInCollection returns whether one and only one document matching the query in the specified collection exists.
func UniqueInCollection(collection *mgo.Collection, selector interface{}) bool {
	n, _ := collection.Find(selector).Count()
	return n == 1
}

func EnsureLocationIndex(collection *mgo.Collection) error {
	index := mgo.Index{
		Key: []string{"$2dsphere:location"},
	}
	return collection.EnsureIndex(index)
}
