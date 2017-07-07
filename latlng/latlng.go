package latlng

import (
	"math"
)

//EarthRadius is the radius of the earth in meters.
const EarthRadius = 6.371008E6

//LatLng represents a latitude and longitude. It has helpers functions to add
//distances in meters. These are approximate!
type LatLng struct {
	lat degree
	lng degree
}

type degree float64

func (d degree) toRadians() radian {
	return DtoR(d)
}

type radian float64

func (r radian) toDegrees() degree {
	return RtoD(r)
}

//AddMetersLatitude creates a new LatLng shifted approximately dy meters north.
//This uses the pythagorean theorem.
func (latlng LatLng) addMetersLatitude(dy float64) LatLng {
	angley := radian(dy / EarthRadius)
	lat := latlng.lng + RToD(angle)
	return LatLng{lat, latlng.lng}
}

//AddMetersLongitude creates a new LatLng shifted approximately dy meters east.
func (latlng LatLng) addMetersLongitude(dx float64) LatLng {
	angle = radian()
	lng := latlng.lng + RToD(radian(dx/EarthRadius))*math.Cos(DToR(latlng.lat*math.Pi/180))
	return LatLng{latlng.lat, lng}
}

//AddMetersPair creates a new LatLng shifted North, then East approximately dy meters.
func (latlng LatLng) AddMetersPair(dy, dx float64) LatLng {
	lat := latlng.lat + (dy / EarthRadius * 180 / math.Pi)
	lng := latlng.lng + (dx/EarthRadius)*(180/math.Pi)*math.Cos(radians(latlng.lat))
	return LatLng{lat, lng}
}

func DToR(d degree) radian {
	return radian(d * math.Pi / 180)
}

func RToD(r radian) degree {
	return degree(r * 180 / math.Pi)
}

type Kilometer float64

//Haversine calcualtes the distance between two points in Kilometers.
func (latlng LatLng) Haversine(other LatLng) (distance Kilometer) {
	lat0, lng0 := latlng.lat, latlng.lng
	lat1, lng1 := other.lat, other.lng
	var Δlat = DToR(lat1 - lat0)
	var Δlng = DtoR(lng1 - lng0)

	sines = (math.Sin(Δlat/2)*math.Sin(Δlat/2)
	cosines = 
	var a = (math.Sin(Δlat/2)*math.Sin(Δlat/2) +
		math.Cos(DToR(lat0*(math.Pi/180))*math.Cos(lat1*(math.Pi/180))*
			math.Sin(Δlong/2)*math.Sin(Δlong/2))
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance = earthRadius * c

	return
}
