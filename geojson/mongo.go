package geojson

import (
	"googlemaps.github.io/maps"
)

type M map[string]interface{}

func WithinPolygonRequest(poly *Polygon) M {
	return M{
		"$geoWithin": M{
			"$geometry": poly.GeoJSON(),
		},
	}
}

func PointFromFactual(record M) Point {
	lat, _ := record["latitutde"].(float64)
	lng, _ := record["longitude"].(float64)
	return NewPoint(lat, lng)
}

func PointFromGoogle(details maps.PlaceDetailsResult) Point {
	location := details.Geometry.Location
	return NewPoint(location.Lat, location.Lng)
}
