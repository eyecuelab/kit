package address

import (
	"strings"

	"github.com/eyecuelab/kit/stringslice"
	"gopkg.in/mgo.v2/bson"
)

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

//Address represents a physical location
type Address struct {
	Street, Extension, POBox string
	Locality                 string //city
	Region                   string //state
	PostalCode               string //zip
	Country                  string
}

func (a *Address) String() string {
	//PO BOX deliberately ignored for now; 9/7/2017
	components := []string{
		a.Street, a.Extension, a.POBox,
		a.Locality, a.Region, a.PostalCode, a.Country,
	}
	return strings.Join(stringslice.NonEmpty(components), ", ")
}
