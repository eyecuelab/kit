package address

import (
	"googlemaps.github.io/maps"
)

type componentType string

const (
	StreetNumber             componentType = "street_number"
	Route                    componentType = "route"
	Subpremise               componentType = "subpremise"
	AdministrativeAreaLevel1 componentType = "administrative_area_level_1"
	Country                  componentType = "country"
	PostalCode               componentType = "postal_code"
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
			switch label {
			case "street_number":
				street.number = val
			case "route":
				street.name = val
			case "subpremise":
				address.Extension = val
			case "administrative_area_level_1":
				address.Region = val
			case "country":
				address.Country = val
			case "postal_code":
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