package jsontaggedbson

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

type Intermediate interface{}

//TaggedIntermediate creates BSON-tagged struct ready for serializing from a JSON struct.
func TaggedIntermediate(obj interface{}) (bdoc Intermediate, err error) {
	JSON, err := json.Marshal(obj)
	if err != nil {
		return bdoc, fmt.Errorf("JSONTaggedBSON: json.Marshal: %v", err)
	}
	err = bson.UnmarshalJSON(JSON, &bdoc)
	if err != nil {
		return bdoc, fmt.Errorf("JSONTaggedBSON: bson.MarshalJSON: %v", err)
	}
	return bdoc, nil
}

//Marshal a JSON-tagged object to BSON.
func Marshal(v interface{}) (BSON []byte, err error) {
	var bdoc Intermediate
	if bdoc, err = TaggedIntermediate(v); err != nil {
		return BSON, fmt.Errorf("TaggedInterface: %v", err)
	}
	return bson.Marshal(bdoc)
}

func isPtr(i interface{}) bool {
	return reflect.ValueOf(i).Kind() == reflect.Ptr
}

//FromTaggedIntermediate assigns the elements of bdoc to the JSON-tagged struct pointed to by v.
func FromTaggedIntermediate(bdoc Intermediate, v interface{}) (err error) {
	if !isPtr(v) {
		return fmt.Errorf("v must be a pointer to a struct")
	}
	var JSON []byte
	if JSON, err = bson.MarshalJSON(bdoc); err != nil {
		return fmt.Errorf("bson.MarshalJSON: %v", err)
	}
	if err = json.Unmarshal(JSON, v); err != nil {
		return fmt.Errorf("json.Unmarshal: %v", err)
	}
	return nil
}

//Unmarshal a BSON-serialized object to the json-tagged object pointed to by v.
func Unmarshal(data []byte, v interface{}) (err error) {
	var bdoc interface{}
	var JSON []byte
	if err = bson.Unmarshal(data, &bdoc); err != nil {
		return fmt.Errorf("bson.Unmarshal: %v", err)
	}
	if JSON, err = bson.MarshalJSON(bdoc); err != nil {
		return fmt.Errorf("bson.MarshalJSON: %v", err)
	}
	if err = json.Unmarshal(JSON, v); err != nil {
		return fmt.Errorf("json.Unmarshal: %v", err)
	}
	return nil
}
