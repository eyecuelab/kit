package geojson

import "gopkg.in/mgo.v2/bson"

type Point [2]float64

func (p *Point) Coordinates() interface{} {
	return *p
}

func (p *Point) Type() Type {
	return PointType
}

func (p *Point) ToBSON() bson.M {
	return toBSON(p)
}
