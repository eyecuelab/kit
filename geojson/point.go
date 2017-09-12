package geojson

type Point struct {
	Lat float64
	Lng float64
}

func (p *Point) Coordinates() Coordinates {
	return [2]float64{p.Lng, p.Lat}
}

func (p *Point) Type() Type {
	return PointType
}

func (p *Point) GeoJSON() GeoJSON {
	return geoJSON(p)
}
