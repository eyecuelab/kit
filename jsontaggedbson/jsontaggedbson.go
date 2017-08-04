package jsontaggedbson

import (
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//TaggedInterface creates BSON-tagged struct ready for serializing from a JSON struct.
func TaggedInterface(obj interface{}) (bdoc interface{}, err error) {
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
	var intermediate interface{}
	if intermediate, err = TaggedInterface(v); err != nil {
		return BSON, fmt.Errorf("TaggedInterface: %v", err)
	}
	return bson.Marshal(intermediate)
}

//Unmarshal a BSON-serialized object to the json-tagged object pointed to by v.
func Unmarshal(data []byte, v interface{}) (err error) {
	var intermediate interface{}
	var JSON []byte
	if err = bson.Unmarshal(data, &intermediate); err != nil {
		return fmt.Errorf("bson.Unmarshal: %v", err)
	}
	if JSON, err = bson.MarshalJSON(intermediate); err != nil {
		return fmt.Errorf("bson.MarshalJSON: %v", err)
	}
	if err = json.Unmarshal(JSON, v); err != nil {
		return fmt.Errorf("json.Unmarshal: %v", err)
	}
	return nil
}
