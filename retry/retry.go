package retry

import (
	"time"
)

//ExpSleep sleeps for 50<<n milliseconds.
func ExpSleep(n uint) {
	time.Sleep(time.Millisecond * 50 << n)
}
