package config

import (
	"reflect"
	"testing"
)

func TestRequiredString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequiredString(tt.args.key); got != tt.want {
				t.Errorf("RequiredString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequiredInt(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequiredInt(tt.args.key); got != tt.want {
				t.Errorf("RequiredInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequiredFloat64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequiredFloat64(tt.args.key); got != tt.want {
				t.Errorf("RequiredFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequiredStringSlice(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequiredStringSlice(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequiredStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFatalCheck(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FatalCheck(tt.args.key)
		})
	}
}

func TestRequiredEnv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequiredEnv(tt.args.key); got != tt.want {
				t.Errorf("RequiredEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
