package geojson

type Polygon []Point

func (poly Polygon) Coordinates() Coordinates {
	coords := make([]interface{}, len(poly))
	for i, p := range poly {
		coords[i] = p.Coordinates()
	}
	return coords
}

func (poly Polygon) Type() Type {
	return PolygonType
}

func (poly Polygon) GeoJSON() GeoJSON {
	return geoJSON(poly)
}

func (poly Polygon) Valid() bool {
	return len(poly) >= 3 && poly[0] == poly[len(poly)-1]
}
