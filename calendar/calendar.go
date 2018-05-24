package calendar

import (
	"time"

	"github.com/rickar/cal"
)

// USCalendar ...
type USCalendar struct {
	Calendar *cal.Calendar
}

// AddWorkHours add number of working hours to a provided time
// excludes weekends, holidays and time between 6pm and 8am
func (c *USCalendar) AddWorkHours(t time.Time, hours int) time.Time {
	if hours < 1 {
		return t
	}

	newT := t.Add(time.Duration(1) * time.Hour)
	hour := newT.Hour()
	if c.Calendar.IsWorkday(newT) && hour >= 8 && hour < 18 {
		hours--
	}

	return c.AddWorkHours(newT, hours)
}

// SetupUSCalendar ...
func SetupUSCalendar() *USCalendar {
	c := &USCalendar{Calendar: cal.NewCalendar()}
	c.Calendar.AddHoliday(
		cal.USIndependence,
		cal.USThanksgiving,
		cal.USChristmas,
		cal.USColumbus,
		cal.USLabor,
		cal.USMLK,
		cal.USMemorial,
		cal.USNewYear,
	)

	return c
}
