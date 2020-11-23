package controllers

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	tables := []struct {
		nowHour   int
		bedtimeAt int
		wakeupAt  int
		sleeping  bool
	}{
		{nowHour: 20, bedtimeAt: 21, wakeupAt: 10, sleeping: false},
		{nowHour: 22, bedtimeAt: 21, wakeupAt: 10, sleeping: true},
		{nowHour: 23, bedtimeAt: 21, wakeupAt: 10, sleeping: true},
		{nowHour: 0, bedtimeAt: 21, wakeupAt: 10, sleeping: true},
		{nowHour: 1, bedtimeAt: 21, wakeupAt: 10, sleeping: true},
		{nowHour: 9, bedtimeAt: 21, wakeupAt: 10, sleeping: true},
		{nowHour: 10, bedtimeAt: 21, wakeupAt: 10, sleeping: false},
		{nowHour: 11, bedtimeAt: 21, wakeupAt: 10, sleeping: false},
	}

	for _, table := range tables {
		st := sleepTime{
			now:       time.Date(2020, 01, 10, table.nowHour, 0, 0, 0, time.Local),
			bedtimeAt: table.bedtimeAt,
			wakeupAt:  table.wakeupAt,
		}

		if st.isSleepTime() != table.sleeping {
			t.Errorf(
				"You have jet-lag. Ha?"+
					": now: %d, bedtime: %d, wakeup: %d",
				table.nowHour, table.bedtimeAt, table.wakeupAt,
			)
		}
	}
}
