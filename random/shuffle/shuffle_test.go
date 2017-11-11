package shuffle

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eyecuelab/kit/random"
)

func TestBytes(t *testing.T) {
	start, _ := random.RandomBytes(10)
	seen := make([][]byte, 0, 100)
	for i := 0; i < 100; i++ {
		perm, _ := Bytes(start)
		for _, prev := range seen {
			assert.NotEqual(t, perm, prev)
		}
		seen = append(seen, perm)
	}
}

func TestInt64s(t *testing.T) {
	start, _ := random.Int64s(10)
	seen := make([][]int64, 0, 100)
	for i := 0; i < 100; i++ {
		perm, _ := Int64s(start)
		for _, prev := range seen {
			assert.NotEqual(t, perm, prev)
		}
		seen = append(seen, perm)
	}
}

func TestFloat64s(t *testing.T) {

}

func TestStrings(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Strings(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("Strings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}
