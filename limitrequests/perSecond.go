package limitrequests

import (
	"sync/atomic"
	"time"
)

//PerSecondLimiter limits the number of requests launches per second. Build with NewTotalLimiter.
//Implements ConcurrencyLimiter.
//Compare with TotalLimiter for limiting the total number of launched requests.
type PerSecondLimiter struct {
	current, Total, completed int64
	requestsThisPeriod        int64
	timeout, pollingPeriod    time.Duration
	maxPerSecond              float64
	beginningOfPollingPeriod  time.Time
}

//NewPerSecondLimiter builds a PerSecondLimiter, an interface for limiting the rate of simulataneous concurrent requests.
func NewPerSecondLimiter(maxPerSecond float64, timeout time.Duration) *PerSecondLimiter {
	return &PerSecondLimiter{
		timeout:       timeout,
		maxPerSecond:  maxPerSecond,
		pollingPeriod: 25 * time.Millisecond,
	}
}

//WaitForSlot waits up to timeout seconds for a slot to open. If one does, it returns true; else, false.
func (psl *PerSecondLimiter) WaitForSlot() bool {
	for wait, i := time.Duration(0), uint(0); wait < psl.timeout; i++ {
		if psl.openSlot() {
			return true
		}
		wait += psl.pollingPeriod
		time.Sleep(psl.pollingPeriod)
	}
	return false
}

//Add a request.
func (psl *PerSecondLimiter) add() {
	atomic.AddInt64(&psl.current, 1)
	atomic.AddInt64(&psl.Total, 1)
	atomic.AddInt64(&psl.requestsThisPeriod, 1)
}

//Done signifies that a request is completed; this usually opens up a slot for another.
func (psl *PerSecondLimiter) Done() {
	atomic.AddInt64(&psl.current, -1)
	atomic.AddInt64(&psl.completed, 1)
}

func (psl *PerSecondLimiter) openSlot() bool {
	if atomic.LoadInt64(&psl.current) == 0 {
		psl.beginningOfPollingPeriod = time.Now() // reset our polling
		return true
	}
	elapsed := time.Now().Sub(psl.beginningOfPollingPeriod).Seconds()
	requestsPerSecond := float64(atomic.LoadInt64(&psl.current)) / elapsed
	return requestsPerSecond < psl.maxPerSecond
}

func (psl *PerSecondLimiter) resetPeriod() {
	atomic.StoreInt64(&psl.requestsThisPeriod, 0)
	psl.beginningOfPollingPeriod = time.Now() // this is not atomic,
	// but it is the only thing that writes to beginningOfPollingPeriod.
}

//WaitAll waits for all current requests to complete. This can wait forever if some requests never finish.
//Use WaitAlltimeout if that's not what you want.
func (psl *PerSecondLimiter) WaitAll() {
	for atomic.LoadInt64(&psl.current) > 0 {
		time.Sleep(psl.pollingPeriod)
	}
}

//WaitAlltimeout waits for timeout duration for all requests to complete. Returns true if no timeout; false if timeout.
// If you want to wait forever, use WaitAll instead.
func (psl *PerSecondLimiter) WaitAlltimeout(timeout time.Duration) bool {
	for wait := time.Duration(0); wait < psl.timeout; wait += psl.pollingPeriod {
		if atomic.LoadInt64(&psl.current) == 0 {
			return true
		}
		time.Sleep(psl.pollingPeriod)
	}
	return false
}
