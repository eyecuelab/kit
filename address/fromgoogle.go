package address

import "googlemaps.github.io/maps"

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
			case "subpremise":
				address.Extension = val
			case "administrative_area_level_1":
				address.Region = val
			case "country":
				address.Country = val
			case "postal_code":
				address.PostalCode = val
			}
			//note: PO box doesn't seem to be handled.
		}
	}
	address.Street = street.number + " " + street.name
	return address
}
