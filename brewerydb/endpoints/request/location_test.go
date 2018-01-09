package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation_Valid(t *testing.T) {
	var loc Location
	assert.Error(t, loc.Valid())
	loc.Region = "CA"
	assert.NoError(t, loc.Valid())
	loc.Page = -1
	assert.Error(t, loc.Valid())
}
