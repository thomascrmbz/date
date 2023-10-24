package date_test

import (
	"testing"

	"github.com/thomascrmbz/date"
)

func TestStartMonth(t *testing.T) {
	t.Logf("Start of this month: %s - %s", date.ThisMonthStart(), date.ThisMonthEnd())
	t.Logf("Start of last month: %s - %s", date.LastMonthStart(), date.LastMonthEnd())
	t.Logf("Start of next month: %s - %s", date.NextMonthStart(), date.NextMonthEnd())
}
