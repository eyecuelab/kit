package structs

type Address struct {
	StreetAddress   string `json:"streetAddress,omitempty"`
	ExtendedAddress string `json:"extendedAddress,omitempty"`
	Locality        string `json:"Locality,omitempty"`
	Region          string `json:"Region,omitempty"`
	PostalCode      string `json:"Postal_code,omitempty"`
}
