package limitrequests

import (
	"sync/atomic"
	"time"
)

//PerSecondMaxLimiter limits the number of requests launches per second, and also hard-caps the total number of requests sent.
//Build with NewPerSecondMaxLimiter.
//Implements ConcurrencyLimiter and MaxLimiter.
//Compare with TotalMaxLimiter for limiting the total number of simultaneous requests, rather than the rate at which they are launched.
type PerSecondMaxLimiter struct {
	*PerSecondLimiter
	maxTotalRequests int64
}

//NewPerSecondMaxLimiter builds a PerSecondMaxLimiter, an interface for limiting the rate of simulataneous concurrent requests.
func NewPerSecondMaxLimiter(maxPerSecond float64, timeout time.Duration, maxTotalRequests int64) *PerSecondMaxLimiter {
	return &PerSecondMaxLimiter{
		NewPerSecondLimiter(maxPerSecond, timeout),
		maxTotalRequests,
	}
}

//HitMaxRequests() returns true if the total requests are
func (psml *PerSecondMaxLimiter) HitMaxRequests() bool {
	return atomic.LoadInt64(&psml.Total) >= atomic.LoadInt64(&psml.maxTotalRequests)
}

//SlotOpenBeforeTimeout returns true if a a slot opens before timeout.
//Masks SlotOpenBeforeTimeOut inherited from PerSecondLimiter
func (psml *PerSecondMaxLimiter) SlotOpenBeforeTimeout() bool {
	if psml.HitMaxRequests() {
		return false
	}

	for wait, i := time.Duration(0), uint(0); wait < psml.timeout; i++ {
		if psml.openSlot() {
			return true
		}
		wait += psml.pollingPeriod
		time.Sleep(psml.pollingPeriod)
	}
	return false
}

//TotalMaxLimiter limits the Total number of simultaneous requests, and also the maximum # of requests launched. Build with NewTotalMaxLimiter. Implements ConcurrencyLimiter.
//Compare with PerSecondLimiter for limiting the rate at which requests are launched.
type TotalMaxLimiter struct {
	*TotalLimiter
	maxTotalRequests int64
}

//NewTotalMaxLimiter builds a new total max limiter.
func NewTotalMaxLimiter(maxSimultaneousRequests, maxTotalRequests int64, timeout time.Duration) *TotalMaxLimiter {
	return &TotalMaxLimiter{
		NewTotalLimiter(maxSimultaneousRequests, timeout),
		maxTotalRequests,
	}
}

//HitMaxRequests returns true if the total requests so far are greater than the max requests.
func (tml *TotalMaxLimiter) HitMaxRequests() bool {
	return atomic.LoadInt64(&tml.Total) >= atomic.LoadInt64(&tml.maxTotalRequests)
}

//SlotOpenBeforeTimeOut returns true if a slot opens before timeout.
func (tml *TotalMaxLimiter) SlotOpenBeforeTimeOut() bool {
	if tml.HitMaxRequests() {
		return false
	}
	for wait, i := time.Duration(0), uint(0); wait < tml.timeout; i++ {
		if tml.openSlot() {
			return true
		}
		wait += tml.pollingPeriod
		time.Sleep(tml.pollingPeriod)
	}
	return false
}
