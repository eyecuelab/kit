package address

import (
	"strings"

	"github.com/eyecuelab/kit/stringslice"
	"googlemaps.github.io/maps"
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

func FromGoogleAddressComponents(components []maps.AddressComponent) (address Address) {
	type routeparts struct {
		name, number string
	}
	var street routeparts
	for _, component := range components {
		val := component.ShortName
		for _, label := range component.Types {
			switch label {
			case "street_number":
				street.name = val
			case "route":
				street.number = val
			case "administrative_area_level_1":
				address.Region = val
			case "country":
				address.Country = val
			case "postal_code":
				address.PostalCode = val
			}
		}
	}
	address.Street = street.number + " " + street.name
	return address
}

//Address represents a physical location
type Address struct {
	Street, Extension, POBox string
	Locality                 string //city
	Region                   string //state
	PostalCode               string //zip
	Country                  string
}

//SharedComponentsOf returns a copy of b, except that where a.component == "", b.component == ""
func (a *Address) SharedComponentsOf(b Address) Address {
	if a.Street == "" {
		b.Street = ""
	}
	if a.Extension == "" {
		b.Extension = ""
	}
	if a.POBox == "" {
		b.POBox = ""
	}
	if a.Locality == "" {
		b.Locality = ""
	}
	if a.Region == "" {
		b.Region = ""
	}
	if a.PostalCode == "" {
		b.PostalCode = ""
	}
	if a.Country == "" {
		b.Country = ""
	}
	return b
}

func (a *Address) String() string {
	//PO BOX deliberately ignored for now; 9/7/2017
	components := []string{
		a.Street, a.Extension, a.POBox,
		a.Locality, a.Region, a.PostalCode, a.Country,
	}
	return strings.Join(stringslice.NonEmpty(components), ", ")
}
