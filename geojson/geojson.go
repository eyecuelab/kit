//Package geojson contains an interface and various structs for 2d and 3d geometry that correspond to the [geojson](http://geojson.org/) format.
package geojson

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

const (
	PointType   = "Point"
	PolygonType = "Polygon"
)

type Geometry interface {
	bson.Getter
	Type() string
	Coordinates() string
	GeoJSON() []byte
}

func geoJSON(geo Geometry) []byte {
	return []byte(fmt.Sprintf(`{"type": %s, "coordinates": %s}`, geo.Type(), geo.Coordinates()))
}
