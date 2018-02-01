package address

import "googlemaps.github.io/maps"

//FromGoogleAddressComponents creates an Addresss from a slice of components, using the AddressCompoment.Types to discriminate.
func FromGoogleAddressComponents(components []maps.AddressComponent) (address Address) {
	var street struct{ name, number string }
	for _, component := range components {
		val := component.ShortName
		for _, label := range component.Types {
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
