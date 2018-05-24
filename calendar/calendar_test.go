package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddWorkHours(t *testing.T) {
	c := SetupUSCalendar()

	losAngeles, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		t.Errorf("Error loading location: %s", err.Error())
	}

	tm, err := time.ParseInLocation(
		"Jan 2, 2006 at 3:04pm (MST)",
		"May 24, 2018 at 11:00am (PDT)", losAngeles)
	if err != nil {
		t.Errorf("Error parsing time: %s", err.Error())
	}

	// same day
	tm2 := c.AddWorkHours(tm, 6)
	assert.Equal(t, 24, tm2.Day())
	assert.Equal(t, 17, tm2.Hour())

	// next day
	tm2 = c.AddWorkHours(tm, 10)
	assert.Equal(t, 25, tm2.Day())
	assert.Equal(t, 11, tm2.Hour())

	// with weekend and labor day
	tm2 = c.AddWorkHours(tm, 20)
	assert.Equal(t, 29, tm2.Day())
	assert.Equal(t, 11, tm2.Hour())
}
