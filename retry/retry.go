package retry

import (
	"time"
)

//ExpSleep sleeps for 50<<n milliseconds.
func ExpSleep(n int) {
	time.Sleep(time.Millisecond * 50 << uint(n))
}
