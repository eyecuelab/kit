package web

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestAPIURL(t *testing.T) {
	viper.Set("root_url", "http://example.com/")
	url := APIURL("resources/1")
	assert.Equal(t, "http://example.com/resources/1", url)
}
