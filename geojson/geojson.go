package geojson

import (
	"fmt"
)

const (
	PointType   = "Point"
	PolygonType = "Polygon"
)

type Geometry interface {
	Type() string
	Coordinates() string
	GeoJSON() []byte
	GetBSON() interface{}
}

func geoJSON(geo Geometry) []byte {
	return []byte(fmt.Sprintf(`{"type": %s, "coordinates": %s}`, geo.Type(), geo.Coordinates()))
}
