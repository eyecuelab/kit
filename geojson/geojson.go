package geojson

import "gopkg.in/mgo.v2/bson"

type Type string

const (
	PointType      = Type("Point")
	PolyGonType    = Type("Polygon")
	MultiPointType = Type("MultiPoint")
)

type GeoJSON interface {
	Type() Type
	Coordinates() interface{}
	ToBSON() bson.M
}

func toBSON(geoJSON GeoJSON) bson.M {
	return bson.M{
		"type":        geoJSON.Type(),
		"coordinates": geoJSON.Coordinates(),
	}
}

