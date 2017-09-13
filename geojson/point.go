package geojson

import "fmt"

type Point struct {
	Lat float64
	Lng float64
}

//Coordinates is the coordinates.
func (p *Point) Coordinates() string {
	return fmt.Sprintf(`[%f, %f]`, p.Lat, p.Lng)
}

//Type is the type.
func (p *Point) Type() string {
	return PointType
}

//GeoJSON formats to GeoJSON:  {type : "Point", coordinates : [lat, lng]}
func (p *Point) GeoJSON() []byte {
	return geoJSON(p)
}
