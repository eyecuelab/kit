package address

import (
	"reflect"
	"testing"
)

func TestNormalizedSharedComponentDistanceSlice(t *testing.T) {
	type args struct {
		placeA Address
		placeB Address
	}
	tests := []struct {
		name          string
		args          args
		wantDistances []int
	}{
		{
			name: "all shared",
			args: args{
				placeA: Address{Street: "foo"},
				placeB: Address{Street: "bar"},
			},
			wantDistances: []int{3},
		},
		{
			name: "some shared",
			args: args{
				placeA: Address{Street: "foo", POBox: "foo"},
				placeB: Address{Street: "bar", Country: "bar"},
			},
			wantDistances: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistances := NormalizedSharedComponentDistanceSlice(tt.args.placeA, tt.args.placeB); !reflect.DeepEqual(gotDistances, tt.wantDistances) {
				t.Errorf("NormalizedSharedComponentDistanceSlice() = %v, want %v", gotDistances, tt.wantDistances)
			}
		})
	}
}
