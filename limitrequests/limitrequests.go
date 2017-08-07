//Package limitrequests contains tools for limiting the number of simulatenously launched requests.
package limitrequests

import (
	"sync/atomic"
	"time"
)

//ConcurrencyLimiter is an interface for limiting the number of simulataneous concurrent requests.
type ConcurrencyLimiter interface {
	add()
	Done()
	WaitForSlot() bool
	WaitAll() //equivalent to sync.WaitGroup.Wait()
	openSlot() bool
	Start() bool
	Total()
}

//Limiter is a concrete struct which implements the values & methods of ConcurrencyLimiter that should be shared across all implementations.
type Limiter struct {
	ConcurrencyLimiter
	CurrentRequests, TotalRequests *int64
	Timeout, PollingPeriod         time.Duration
}

//Start waits Timeout seconds for a for a slot to open. If a slot opens, it adds 1 to the current requests and returns true.
//otherwise, it returns false.
func (l *Limiter) Start() bool {
	if timeout := !l.WaitForSlot(); timeout {
		return false
	}

	l.add()
	return true
}

//Add a request.
func (l *Limiter) add() {
	atomic.AddInt64(l.CurrentRequests, 1)
	atomic.AddInt64(l.TotalRequests, 1)
}

//WaitForSlot returns true if a slot opens for requests before Timeout.
func (l *Limiter) WaitForSlot() bool {
	for wait, i := time.Duration(0), uint(0); wait < l.Timeout; i++ {
		if l.openSlot() {
			return true
		}
		wait += l.PollingPeriod
		time.Sleep(l.PollingPeriod)
	}
	return false
}

//WaitAll waits for all current requests to complete. This can wait forever if some requests never finish.
//Use WaitAllTimeOut if that's not what you want.
func (l *Limiter) WaitAll() {
	for atomic.LoadInt64(l.CurrentRequests) > 0 {
		time.Sleep(l.PollingPeriod)
	}
}

//WaitAllTimeout waits for timeout duration for all requests to complete. If you want to wait forever, use WaitAll instead.
func (l *Limiter) WaitAllTimeout(timeout time.Duration) bool {
	for wait := time.Duration(0); wait < l.Timeout; wait += l.PollingPeriod {
		if atomic.LoadInt64(l.CurrentRequests) == 0 {
			return true
		}
		time.Sleep(l.PollingPeriod)
	}
	return false
}

//TotalLimiter limits the total number of current requests. Build with NewTotalLimiter.
type TotalLimiter struct {
	*Limiter
	MaxRequests int64 `json:"max_requests,omitempty"`
}

func (psl *TotalLimiter) openSlot() bool {
	return atomic.LoadInt64(psl.CurrentRequests) < psl.MaxRequests
}

//NewTotalLimiter builds a TotalLimiter
func NewTotalLimiter(maxRequests int64, timeOut time.Duration) *TotalLimiter {
	limiter := Limiter{Timeout: timeOut}
	return &TotalLimiter{&limiter, maxRequests}
}

//PerSecondLimiter limits the number of requests to an amount per second. Build with NewPerSecondLimiter
type PerSecondLimiter struct {
	*Limiter
	MaxRequestsPerSecond        float64 `json:"max_requests_per_second,omitempty"`
	startTimeOfLastRequestCycle time.Time
}

func (psl *PerSecondLimiter) openSlot() bool {
	if atomic.LoadInt64(psl.CurrentRequests) == 0 {
		// no live requests, so we need to reset the time from which we're counting requests per-second.
		psl.startTimeOfLastRequestCycle = time.Now()
		return true
	}
	return psl.requestsPerSecond() > psl.MaxRequestsPerSecond
}

func (psl *PerSecondLimiter) requestsPerSecond() float64 {
	elapsed := time.Now().Sub(psl.startTimeOfLastRequestCycle).Seconds()
	return elapsed / float64(atomic.LoadInt64(psl.CurrentRequests))
}

//NewPerSecondLimiter builds a PerSecondLimiter.
func NewPerSecondLimiter(maxRequestsPerSecond float64, timeOut time.Duration) *PerSecondLimiter {
	limiter := Limiter{Timeout: timeOut}
	return &PerSecondLimiter{&limiter, maxRequestsPerSecond, time.Now()}
}
