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

func (poly Polygon) GetBSON() (interface{}, error) {
	coords := make([][2]float64, len(poly))
	for i, p := range poly {
		coords[i] = [2]float64{p.Lng(), p.Lat()}
	}
	bdoc := map[string]interface{}{
		"type":        poly.Type(),
		"coordinates": coords,
	}
	return bdoc, nil
}
