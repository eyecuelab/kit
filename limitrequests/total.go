//Package limitrequests contains tools for limiting the number of simulatenously launched requests.
package limitrequests

import (
	"sync/atomic"
	"time"
)

//TotalLimiter limits the Total number of simultaneous requests. Build with NewTotalLimiter. Implements ConcurrencyLimiter.
//Compare with PerSecondLimiter for limiting the rate at which requests are launched.
type TotalLimiter struct {
	current, Total, completed int64
	timeout, pollingPeriod    time.Duration
	maxRequests               int64
}

//Start waits timeout seconds for a for a slot to open. If a slot opens, it adds 1 to the current requests and returns true.
//otherwise, it returns false.
func (tl *TotalLimiter) Start() bool {
	if tl.SlotOpenBeforeTimeOut() {
		tl.add()
		return true
	}

	return true

}

//Add a request.
func (tl *TotalLimiter) add() {
	atomic.AddInt64(&tl.current, 1)
	atomic.AddInt64(&tl.Total, 1)
}

//Done signifies that a request is completed; this usually opens up a slot for another.
func (tl *TotalLimiter) Done() {
	atomic.AddInt64(&tl.current, -1)
	atomic.AddInt64(&tl.completed, 1)
}

//SlotOpenBeforeTimeOut returns true if a slot opens before timeout.
func (tl *TotalLimiter) SlotOpenBeforeTimeOut() bool {
	for wait, i := time.Duration(0), uint(0); wait < tl.timeout; i++ {
		if tl.openSlot() {
			return true
		}
		wait += tl.pollingPeriod
		time.Sleep(tl.pollingPeriod)
	}
	return false
}

//Wait waits for all current requests to complete. This can wait forever if some requests never finish.
//Analagous to sync.WaitGroup.Wait()
func (tl *TotalLimiter) Wait() {
	for atomic.LoadInt64(&tl.current) > 0 {
		time.Sleep(tl.pollingPeriod)
	}
}

//WaitTimeout waits up to timeout duration for all current requests to complete. It returns true if
//all requests completed in that time.
func (tl *TotalLimiter) WaitTimeout(timeout time.Duration) bool {
	for wait := time.Duration(0); wait < tl.timeout; wait += tl.pollingPeriod {
		if atomic.LoadInt64(&tl.current) == 0 {
			return true
		}
		time.Sleep(tl.pollingPeriod)
	}
	return false
}

func (tl *TotalLimiter) openSlot() bool {
	return atomic.LoadInt64(&tl.current) < tl.maxRequests
}

//NewTotalLimiter builds a TotalLimiter, an interface for limiting the number of simulataneous concurrent requests.
func NewTotalLimiter(maxSimultaneousRequests int64, timeout time.Duration) *TotalLimiter {
	return &TotalLimiter{
		timeout:       timeout,
		maxRequests:   maxSimultaneousRequests,
		pollingPeriod: 25 * time.Millisecond,
	}
}
