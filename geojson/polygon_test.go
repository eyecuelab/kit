package geojson

import (
	"fmt"
	"reflect"
	"testing"
)

var poly = Polygon{
	{-1, -1}, {-1, 1}, {1, 1}, {-1, -1},
}

func TestPolygon_GeoJSON(t *testing.T) {
	tests := []struct {
		name string
		poly Polygon
		want []byte
	}{
		{"trivialPolygon", poly, []byte(fmt.Sprintf(`{"type": %s, "coordinates": %s}`, poly.Type(), wantCoords))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.poly.GeoJSON(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Polygon.GeoJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

var wantCoords = `[[-1.000000, -1.000000], [-1.000000, 1.000000], [1.000000, 1.000000], [-1.000000, -1.000000]]`

func TestPolygon_Coordinates(t *testing.T) {
	tests := []struct {
		name string
		poly Polygon
		want string
	}{
		{"trivial", poly, wantCoords},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.poly.Coordinates(); got != tt.want {
				t.Errorf("Polygon.Coordinates() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
