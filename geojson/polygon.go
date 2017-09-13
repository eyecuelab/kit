package geojson

import "strings"

type Polygon []Point

func (poly Polygon) Coordinates() string {
	coords := make([]string, len(poly))
	for i, p := range poly {
		coords[i] = p.Coordinates()
	}
	return "[" + strings.Join(coords, ", ") + "]"
}

func (poly Polygon) Type() string {
	return PolygonType
}

func (poly Polygon) GeoJSON() []byte {
	return geoJSON(poly)
}
