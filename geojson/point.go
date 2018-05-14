package geojson

import (
	"fmt"
)

type Point [2]float64

//GeoJson stores points as [lng, lat], everywhere else we use lat, lng
func NewPoint(lat float64, lng float64) Point {
	return Point{lng, lat}
}

func (p Point) Lat() float64 {
	return p[1]
}

func (p Point) Lng() float64 {
	return p[0]
}

//Coordinates is the coordinates.
func (p *Point) Coordinates() string {
	return fmt.Sprintf(`[%f, %f]`, p.Lat(), p.Lng())
}

//Type is the type.
func (p *Point) Type() string {
	return PointType
}

//GeoJSON formats to GeoJSON:  {type : "point", coordinates : [lat, lng]}
func (p *Point) GeoJSON() []byte {
	return geoJSON(p)
}

func (p *Point) String() string {
	return fmt.Sprintf("%f, %f", p.Lat(), p.Lng())
}

func (p *Point) GetBSON() (interface{}, error) {
	bdoc := map[string]interface{}{
		"type":        p.Type(),
		"coordinates": p,
	}
	return bdoc, nil
}
