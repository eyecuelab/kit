package mongotest

import (
	"sync/atomic"
	"github.com/globalsign/mgo"
)

var namespacer int64

func SetupTestDB(name string) *mgo.Database {
	session, _ := mgo.Dial("mongodb://localhost/" + name)
	db := session.DB("test")
	db.DropDatabase()
	return db
}

type TestDB struct {
	*mgo.Database
}

func (db TestDB) TestCollection(v []interface{}) *mgo.Collection {
	name := string(atomic.AddInt64(&namespacer, 1))
	c := db.C(name)
	c.Insert(v...)
	return c
}
