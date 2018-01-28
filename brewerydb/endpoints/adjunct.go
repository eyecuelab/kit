package endpoints

type Beer struct {
	ID                        string `json:"id,omitempty"`
	Name                      string `json:"name,omitempty"`
	Description               string `json:"description,omitempty"`
	FoodPairings              string `json:"foodPairings,omitempty"`
	OriginalGravity           string `json:"originalGravity,omitempty"`
	ABV                       string `json:"abv,omitempty"`
	GlasswareID               string `json:"glassware_id,omitempty"`
	Glass                     string `json:"glass,omitempty"`
	StyleID                   string `json:"style_id,omitempty"`
	Style                     string `json:"style,omitempty"`
	IsOrganic                 string `json:"is_organic,omitempty"`
	Labels                    string `json:"labels,omitempty"`
	ServingTemperature        string `json:"serving_temperature,omitempty"`
	ServingTemperatureDisplay string `json:"serving_temperature_display,omitempty"`
	Status                    string `json:"status,omitempty"`
	StatusDisplay             string `json:"status_display,omitempty"`
	AvailableID               string `json:"available_id,omitempty"`
	Available                 string `json:"available,omitempty"`
	BeerVariationID           string `json:"beer_variation_id,omitempty"`
	BeerVariation             string `json:"beer_variation,omitempty"`
	Years                     string `json:"years,omitempty"`
}
