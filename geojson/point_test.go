package geojson

import (
	"fmt"
	"reflect"
	"testing"
)

var p = Point{1, 2}

func TestPoint_GeoJSON(t *testing.T) {
	tests := []struct {
		name string
		p    *Point
		want []byte
	}{
		{"trivial", &p, []byte(`{"type": Point, "coordinates": [1.000000, 2.000000]}`)},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GeoJSON(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.GeoJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Coordinates(t *testing.T) {
	tests := []struct {
		name string
		p    *Point
		want string
	}{
		{"unit", &Point{1, 0}, fmt.Sprintf(`[%f, %f]`, 1., 0.)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Coordinates(); got != tt.want {
				t.Errorf("Point.Coordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
