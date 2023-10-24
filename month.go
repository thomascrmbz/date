package date

import "time"

func MonthStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func MonthEnd(t time.Time) time.Time            { return MonthStart(t).AddDate(0, 1, 0) }
func MonthStartWithOffset(offset int) time.Time { return MonthStart(time.Now()).AddDate(0, offset, 0) }
func MonthEndWithOffset(offset int) time.Time   { return MonthEnd(time.Now()).AddDate(0, offset, 0) }
func ThisMonthStart() time.Time                 { return MonthStartWithOffset(0) }
func ThisMonthEnd() time.Time                   { return MonthEndWithOffset(0) }
func LastMonthStart() time.Time                 { return MonthStartWithOffset(-1) }
func LastMonthEnd() time.Time                   { return MonthEndWithOffset(-1) }
func NextMonthStart() time.Time                 { return MonthStartWithOffset(1) }
func NextMonthEnd() time.Time                   { return MonthEndWithOffset(1) }
