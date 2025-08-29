package time2

import (
	"sync/atomic"
	"time"
)

var diff int64 = 0

// Now returns the current local time.
func Now() time.Time {
	offset := atomic.LoadInt64(&diff)
	if offset == 0 {
		return time.Now()
	}
	return time.Now().Add(time.Duration(offset))
}

// NowP returns the current local time pointer.
func NowP() *time.Time {
	t := Now()
	return &t
}

// SetTime sets the current time
func SetTime(t time.Time) {
	d := t.Sub(time.Now())
	SetDiff(d)
}

// SetDiff 设置时间偏移量
func SetDiff(d time.Duration) {
	atomic.StoreInt64(&diff, int64(d))
}

// GetDiff returns the difference
func GetDiff() time.Duration {
	return time.Duration(diff)
}

// MaximumSleep10 最大睡觉时间
func MaximumSleep10(d time.Duration) {
	if d > time.Second {
		d = time.Second
	}
	if d > 0 {
		time.Sleep(d)
	}
}

// Until returns the duration until t.
// It is shorthand for t.Sub(time.Now()).
func Until(t time.Time) time.Duration {
	return t.Sub(Now())
}
