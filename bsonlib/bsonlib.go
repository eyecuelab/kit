package bsonlib

import (
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//JSONTaggedBSON turns an object into BSON, using it's JSON tags.
func JSONTaggedBSON(obj interface{}) (bdoc interface{}, err error) {
	jdoc, err := json.Marshal(obj)
	if err != nil {
		return bdoc, fmt.Errorf("JSONTaggedBSON: json.Marshal: %v", err)
	}
	bdoc, err = bson.MarshalJSON(jdoc)
	if err != nil {
		return bdoc, fmt.Errorf("JSONTaggedBSON: bson.MarshalJSON: %v", err)
	}
	return bdoc, nil
}
