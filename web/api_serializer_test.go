package web

import (
	"testing"

	"github.com/spf13/viper"
)

func TestAPIURL(t *testing.T) {
	viper.Set("root_url", "http://example.com/")
	url := APIURL("resources/1")
	expected := "http://example.com/resources/1"
	if url != expected {
		t.Errorf("Incorrect APIURL, expected: %s, received: %s.", expected, url)
	}
}
