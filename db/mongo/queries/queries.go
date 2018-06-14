package queries

import (
	"time"
	"github.com/globalsign/mgo/bson"
)

func TimestampUpdate(name string) bson.M {
	return bson.M{
		"$set": bson.M{
			name: time.Now(),
		},
	}
}
