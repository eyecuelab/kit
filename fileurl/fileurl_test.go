package fileurl

import (
	"testing"
)

func TestIsFileURL(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal_website", args{"https://www.google.com"}, false},
		{"long_strange_domain", args{"http://www.foo.bar.baz/x/y.json"}, true},
		{"forgot http", args{"efron.zone/examplejson.bgg"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFileURL(tt.args.path); got != tt.want {
				t.Errorf("IsFileURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
