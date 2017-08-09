//Package limitrequests contains tools for limiting the number of simulatenously launched requests.
package limitrequests

import (
	"sync/atomic"
	"time"
)

//ConcurrencyLimiter is an interface for limiting the number of simulataneous concurrent requests.
type ConcurrencyLimiter interface {
	add()
	openSlot() bool
	Done()
	WaitForSlot() bool
	WaitAll() //equivalent to sync.WaitGroup.Wait()
	Start() bool
}

//Limiter is a concrete struct which implements the values & methods of ConcurrencyLimiter that should be shared across all implementations.
type Limiter struct {
	current, Total, completed int64
	timeout, pollingPeriod    time.Duration
}

//Start waits timeout seconds for a for a slot to open. If a slot opens, it adds 1 to the current requests and returns true.
//otherwise, it returns false.
func (tl *TotalLimiter) Start() bool {
	if timeout := !tl.WaitForSlot(); timeout {
		return false
	}

	tl.add()
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

//WaitForSlot returns true if a slot opens for requests before timeout.
func (tl *TotalLimiter) WaitForSlot() bool {
	for wait, i := time.Duration(0), uint(0); wait < tl.timeout; i++ {
		if tl.openSlot() {
			return true
		}
		wait += tl.pollingPeriod
		time.Sleep(tl.pollingPeriod)
	}
	return false
}

//WaitAll waits for all current requests to complete. This can wait forever if some requests never finish.
//Use WaitAlltimeout if that's not what you want.
func (tl *TotalLimiter) WaitAll() {
	for atomic.LoadInt64(&tl.current) > 0 {
		time.Sleep(tl.pollingPeriod)
	}
}

//WaitAlltimeout waits for timeout duration for all requests to complete. Returns true if no timeout; false if timeout.
// If you want to wait forever, use WaitAll instead.
func (tl *TotalLimiter) WaitAlltimeout(timeout time.Duration) bool {
	for wait := time.Duration(0); wait < tl.timeout; wait += tl.pollingPeriod {
		if atomic.LoadInt64(&tl.current) == 0 {
			return true
		}
		recover()
		time.Sleep(tl.pollingPeriod)
	}
	return false
}

//TotalLimiter limits the Total number of current requests. Build with NewTotalLimiter.
type TotalLimiter struct {
	current, Total, completed int64
	timeout, pollingPeriod    time.Duration
	maxRequests               int64 `json:"max_requests,omitempty"`
}

func (tl *TotalLimiter) openSlot() bool {
	return atomic.LoadInt64(&tl.current) < tl.maxRequests
}

//NewTotalLimiter builds a TotalLimiter, an interface for limiting the number of simulataneous concurrent requests.
func NewTotalLimiter(maxRequests int64, timeout time.Duration) *TotalLimiter {
	return &TotalLimiter{
		timeout:       timeout,
		maxRequests:   maxRequests,
		pollingPeriod: 25 * time.Millisecond,
	}
}
