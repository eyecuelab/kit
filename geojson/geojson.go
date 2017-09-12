package geojson

type Type string

const (
	PointType   = Type("Point")
	PolygonType = Type("Polygon")
)

type Geometry interface {
	Type() Type
	Coordinates() Coordinates
	GeoJSON() GeoJSON
}

type Coordinates interface{}

type GeoJSON map[string]interface{}

func geoJSON(geo Geometry) GeoJSON {
	return GeoJSON{
		"type":        geo.Type(),
		"coordinates": geo.Coordinates(),
	}
}
