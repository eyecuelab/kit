package geojson

import (
	"github.com/kellydunn/golang-geo"
	"gopkg.in/mgo.v2/bson"
)

type Polygon struct {
	*geo.Polygon
}

func (poly *Polygon) Coordinates() Coordinates {
	points := poly.Points()
	coords := make([]interface{}, len(points))
	for i, p := range points {
		point := Point{p}
		coords[i] = point.Coordinates()
	}
	return coords
}

func (poly *Polygon) Type() Type {
	return PolygonType
}

func (poly *Polygon) GeoJSON() GeoJSON {
	return geoJSON(poly)
}

func WithinPolygonRequest(poly *Polygon) bson.M {
	return bson.M{
		"$geoWithin": bson.M{
			"$geometry": poly.GeoJSON(),
		},
	}
}

func NewPolygon(latlngs [][2]float64) Polygon {
	points := make([]*geo.Point, len(latlngs))
	for i, latlng := range latlngs {
		lat, lng := latlng[0], latlng[1]
		points[i] = geo.NewPoint(lat, lng)
	}
	return Polygon{geo.NewPolygon(points)}
}
