package address

import "gopkg.in/mgo.v2/bson"

//FromFactualRecord builds an address from a serialized record from factual
func FromFactualRecord(factualRecord bson.M) Address {
	getStr := func(key string) string {
		val, _ := factualRecord[key]
		str, _ := val.(string)
		return str
	}
	return Address{
		Street:     getStr("address"),
		Extension:  getStr("address_extended"),
		POBox:      getStr("po_box"),
		Locality:   getStr("locality"),
		PostalCode: getStr("postcode"),
		Region:     getStr("region"),
		Country:    getStr("country"),
	}
}
