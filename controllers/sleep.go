package controllers

import "time"

type sleepTime struct {
	now       time.Time
	bedtimeAt int // hour
	wakeupAt  int // hour
}

func (s sleepTime) createTimeWithHour(hour int, day int) time.Time {
	return time.Date(s.now.Year(), s.now.Month(), day, hour, 0, 0, 0, time.Local)
}

func (s sleepTime) isSleepTime() bool {
	day := s.now.Day()
	hour, _, _ := s.now.Clock()
	var bedtimeDay int
	var wakeupDay int
	if s.bedtimeAt < hour && hour <= 23 {
		bedtimeDay = day
		wakeupDay = day + 1
	} else {
		bedtimeDay = day - 1
		wakeupDay = day
	}
	bedtime := s.createTimeWithHour(s.bedtimeAt, bedtimeDay)
	wakeup := s.createTimeWithHour(s.wakeupAt, wakeupDay)
	return s.now.After(bedtime) && s.now.Before(wakeup)
}
