package request

import "time"
import "github.com/eyecuelab/kit/brewerydb/structs"

const (
	errMustSetLocation BadRequestError = "must set one of the following attributes: locality, postalCode, region"
	errBadPage         BadRequestError = "page must be and int >=0"
)

type Location struct {
	structs.Address

	Page int      `json:"p,omitempty"`
	IDs  []string `json:"ids,omitempty"`

	IsPrimary      bool                   `json:"isPrimary,omitempty"`
	InPlanning     bool                   `json:"inPlanning,omitempty"`
	IsClosed       bool                   `json:"isClosed,omitempty"`
	LocationType   []structs.LocationType `json:"locationType,omitempty"`
	CountryIsoCode string                 `json:"countryIsoCode,omitempty"`
	Since          time.Time              `json:"since,omitempty"`
	Status         string                 `json:"status,omitempty"`
}

func (loc Location) Valid() error {
	if loc.Locality == "" && loc.Region == "" && loc.PostalCode == "" {
		return errMustSetLocation
	} else if loc.Page < 0 {
		return errBadPage
	}
	return nil
}

type BadRequestError string

func (err BadRequestError) Error() string { return string(err) }
