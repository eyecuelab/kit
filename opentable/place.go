package opentable

type Place struct {
	RID            string `json:"rid,omitempty" bson:"rid,omitempty"`
	Name           string `json:"name,omitempty" bson:"name,omitempty"`
	Lat            string `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Lng            string `json:"longitude,omitempty" bson:"longitude,omitempty"`
	StreetAddress  string `json:"address,omitempty" bson:"address,omitempty"`
	City           string `json:"city,omitempty" bson:"city,omitempty"`
	State          string `json:"state,omitempty" bson:"state,omitempty"`
	PostalCode     string `json:"postal_code,omitempty" bson:"postal_code,omitempty"`
	Country        string `json:"country,omitempty" bson:"country,omitempty"`
	PhoneNumber    string `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	MetroName      string `json:"metro_name,omitempty" bson:"metro_name,omitempty"`
	ReservationURL string `json:"reservation_url,omitempty" bson:"reservation_url,omitempty"`
	ProfileURL     string `json:"profile_url,omitempty" bson:"profile_url,omitempty"`
	InGroup        bool   `json:"is_restaraunt_in_group,omitempty" bson:"is_restaraunt_in_group,omitempty"`
}
