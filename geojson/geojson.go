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
