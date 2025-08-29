package time2

import "time"

// EndOfDay 获取今天的23:59:59 -本地时间
func EndOfDay(now time.Time) time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
}

// BeginOfDay 获取今天的凌晨 -本地时间
func BeginOfDay(now time.Time) time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Location())
}
