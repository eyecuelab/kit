package psql

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
)

func TestTestDB(t *testing.T) {
	_, err := TestDB()
	assert.NoError(t, err)
}
