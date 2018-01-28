package structs

type Location struct {
	Address `json:"address,omitempty"`
	LatLng  `json:"lat_lng,omitempty"`

	ID                       string       `json:"id,omitempty"`
	Name                     string       `json:"name,omitempty"`
	Phone                    string       `json:"phone,omitempty"`
	IsPrimary                bool         `json:"isPrimary,omitempty"`
	InPlanning               bool         `json:"inPlanning,omitempty"`
	IsClosed                 bool         `json:"isClosed,omitempty"`
	OpenToPublic             bool         `json:"openToPublic,omitempty"`
	HoursOfOperation         interface{}  `json:"hoursOfOperation,omitempty"`
	HoursOfOperationExplicit interface{}  `json:"hoursOfOperation_explicit,omitempty"`
	HoursOfOperationNotes    string       `json:"hoursOfOperation_notes,omitempty"`
	TourInfo                 string       `json:"tourInfo,omitempty"`
	LocationType             LocationType `json:"locationType,omitempty"`
	LocationTypeDisplay      string       `json:"locationType_display,omitempty"`
	CountryIsoCode           string       `json:"countryIsoCode,omitempty"`
	YearOpened               string       `json:"yearOpened,omitempty"`
	BreweryID                string       `json:"breweryId,omitempty"`
	Brewery                  Brewery      `json:"brewery,omitempty"`
}
