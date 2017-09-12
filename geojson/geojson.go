package geojson

type Type string

const (
	PointType      = Type("Point")
	PolygonType    = Type("Polygon")
	MultiPointType = Type("MultiPoint")
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
