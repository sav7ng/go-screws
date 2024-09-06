package calendar

import (
	"time"
)

type CalendarTool struct{}

var TimeZoneLocation = time.FixedZone("GMT", 8*3600)

// FirstDayOfYear returns the first day 00:00:00 of the year in which the specified date resides
func (tool *CalendarTool) FirstDayOfYear(date time.Time) time.Time {
	year := date.Year()
	firstDay := time.Date(year, 1, 1, 0, 0, 0, 0, date.Location())
	return firstDay
}

// LastDayOfYear returns the last day 00:00:00 of the year in which the specified date resides
func (tool *CalendarTool) LastDayOfYear(date time.Time) time.Time {
	year := date.Year()
	lastDay := time.Date(year+1, time.January, 1, 0, 0, 0, 0, date.Location()).AddDate(0, 0, -1)
	return lastDay
}

// FirstDayOfMonth returns the first day 00:00:00 of the month in which the specified date resides
func (tool *CalendarTool) FirstDayOfMonth(date time.Time) time.Time {
	year, month, _ := date.Date()
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, date.Location())
	return firstDay
}

// LastDayOfMonth returns the last day 00:00:00 of the month in which the specified date resides
func (tool *CalendarTool) LastDayOfMonth(date time.Time) time.Time {
	year, month, _ := date.Date()
	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, date.Location())
	return lastDay
}

// FirstDayOfWeek returns the Monday 00:00:00 of the week in which the specified date resides
// the first day of the week is Monday
func (tool *CalendarTool) FirstDayOfWeek(date time.Time) time.Time {
	weekday := int(date.In(TimeZoneLocation).Weekday())
	monday := date.AddDate(0, 0, 1-weekday)
	monday = time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, date.Location())
	return monday
}

// LastDayOfWeek returns the Monday 00:00:00 of the week in which the specified date resides
// the last day of the week is Sunday
func (tool *CalendarTool) LastDayOfWeek(date time.Time) time.Time {
	weekday := int(date.In(TimeZoneLocation).Weekday())
	sunday := date.AddDate(0, 0, 7-weekday)
	sunday = time.Date(sunday.Year(), sunday.Month(), sunday.Day(), 0, 0, 0, 0, date.Location())
	return sunday
}

// SomeDayBefore returns the time 00:00:00 before the days on the date
func (tool *CalendarTool) SomeDayBefore(date time.Time, days int) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location()).AddDate(0, 0, -days)
}

// SomeDayBefore returns the time 00:00:00 after the days on the date
func (tool *CalendarTool) SomeDayAfter(date time.Time, days int) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location()).AddDate(0, 0, days)
}
