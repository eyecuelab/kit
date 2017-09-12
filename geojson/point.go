package geojson

import geo "github.com/kellydunn/golang-geo"

type Point struct {
	*geo.Point
}

func (p *Point) Coordinates() Coordinates {
	return [2]float64{p.Lng(), p.Lat()}
}

func (p *Point) Type() Type {
	return PointType
}

func (p *Point) GeoJSON() GeoJSON {
	return geoJSON(p)
}

// Returns a new Point populated by the passed in latitude (lat) and longitude (lng) values.
func NewPoint(lat, lng float64) Point {
	return Point{geo.NewPoint(lat, lng)}
}
