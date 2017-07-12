package latlng

import (
	"math"
)

//EarthRadius is the radius of the earth in meters.
const EarthRadius = 6.371008E6

//LatLng represents a latitude and longitude. It has helpers functions to add
//distances in meters. These are approximate!
type LatLng struct {
	Lat float64
	Lng float64
}

//DToR converts degrees to Radians. r = dπ/180 Does not bounds check.
func DToR(d float64) float64 {
	return d * math.Pi / 180
}

//RToD converts Radians to Degrees. d=180r/π. Does not bounds check.
func RToD(r float64) float64 {
	return r * 180 / math.Pi
}

//AddMetersLatitude creates a new LatLng shifted approximately dy meters north.
//This uses the pythagorean theorem; which is less accurate than other methods.
func (latlng LatLng) addMetersLatitude(dy float64) LatLng {
	return LatLng{latlng.Lat + RToD(dy/EarthRadius), latlng.Lng}
}

//AddMetersLongitude creates a new LatLng shifted approximately dy meters east.
func (latlng LatLng) addMetersLongitude(dx float64) LatLng {
	return LatLng{latlng.Lat, latlng.Lng + RToD(dx/EarthRadius)*math.Cos(DToR(latlng.Lat))}
}

//AddMetersPair creates a new LatLng shifted North, then East approximately dy meters.
func (latlng LatLng) AddMetersPair(dy, dx float64) LatLng {
	return LatLng{Lat: latlng.Lat + RToD(dy/EarthRadius),
		Lng: latlng.Lng + RToD(dx/EarthRadius*math.Cos(DToR(latlng.Lat)))}
}

//Haversine returns the great-circle distance in meters between two points on earth.
func Haversine(q, r LatLng) float64 {
	Δlat := DToR(r.Lat - q.Lat)
	Δlng := DToR(r.Lng - q.Lng)
	a := (math.Sin(Δlat/2) * math.Sin(Δlat/2)) +
		math.Cos(DToR(q.Lat)*math.Cos(DToR(r.Lat)))*
			math.Sin(Δlng/2)*math.Sin(Δlng/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EarthRadius * c
}
