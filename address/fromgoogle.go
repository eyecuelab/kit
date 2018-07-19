package address

import (
	"googlemaps.github.io/maps"
)

type componentType string

const (
	StreetNumber             componentType = "street_number"
	Route                    componentType = "route"                       //street
	Subpremise               componentType = "subpremise"                  //unit number
	Locality                 componentType = "Locality"                    //city
	AdministrativeAreaLevel1 componentType = "administrative_area_level_1" //state
	PostalCode               componentType = "postal_code"                 //zip
	Country                  componentType = "country"
)

//FromGoogleAddressComponents creates an Addresss from a slice of components, using the AddressCompoment.Types to discriminate.
func FromGoogleAddressComponents(addressComponents []maps.AddressComponent, whitelist ...componentType) (address Address) {
	var street struct{ name, number string }
	for _, component := range addressComponents {
		val := component.ShortName
		for _, label := range component.Types {
			if !isWhitelisted(whitelist, label) {
				continue
			}
			switch componentType(label) {
			case StreetNumber:
				street.number = val
			case Route:
				street.name = val
			case Subpremise:
				address.Extension = val
			case Locality:
				address.Locality = val
			case AdministrativeAreaLevel1:
				address.Region = val
			case Country:
				address.Country = val
			case PostalCode:
				address.PostalCode = val
			default: //pass
			}
			//note: PO box doesn't seem to be handled.
		}
	}

	if street.name != "" {
		if street.number != "" {
			address.Street = street.number + " " + street.name
		} else {
			address.Street = street.name
		}
	}

	return address
}

func isWhitelisted(whitelist []componentType, componentLabel string) bool {
	if len(whitelist) == 0 {
		return true
	}

	for _, wl := range whitelist {
		if string(wl) == componentLabel {
			return true
		}
	}

	return false
}
