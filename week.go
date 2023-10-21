package date

import (
	"time"
)

func WeekStart(t time.Time) time.Time {
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, -1)
	}
	return t
}

func WeekEnd(t time.Time) time.Time {
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	for t.Weekday() != time.Sunday {
		t = t.AddDate(0, 0, 1)
	}
	return t.AddDate(0, 0, 1)
}

func WeekStartWithOffset(offset int) time.Time { return WeekStart(time.Now()).AddDate(0, 0, 7*offset) }
func WeekEndWithOffset(offset int) time.Time   { return WeekEnd(time.Now()).AddDate(0, 0, 7*offset) }
func ThisWeekStart() time.Time                 { return WeekStartWithOffset(0) }
func ThisWeekEnd() time.Time                   { return WeekEndWithOffset(0) }
func LastWeekStart() time.Time                 { return WeekStartWithOffset(-1) }
func LastWeekEnd() time.Time                   { return WeekEndWithOffset(-1) }
func NextWeekStart() time.Time                 { return WeekStartWithOffset(1) }
func NextWeekEnd() time.Time                   { return WeekEndWithOffset(1) }
