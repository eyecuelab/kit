package limitrequests

import (
	"sync"
	"sync/atomic"
	"time"
)

//PerSecondLimiter limits the number of requests launches per second. Build with NewPerSecondLimiters.
//Implements ConcurrencyLimiter.
//Compare with TotalLimiter for limiting the total number of launched requests.
type PerSecondLimiter struct {
	current, Total, completed int64
	requestsThisPeriod        int64
	timeout, pollingPeriod    time.Duration
	maxPerSecond              float64

	startOfFirstRequest time.Time
	mutex               sync.Mutex //guards startOfFirstRequest
}

//NewPerSecondLimiter builds a PerSecondLimiter, an interface for limiting the rate of simulataneous concurrent requests.
func NewPerSecondLimiter(maxPerSecond float64, timeout time.Duration) *PerSecondLimiter {
	return &PerSecondLimiter{
		timeout:       timeout,
		maxPerSecond:  maxPerSecond,
		pollingPeriod: 25 * time.Millisecond,
	}
}

//SlotOpenBeforeTimeout returns true if a a slot opens before timeout
func (psl *PerSecondLimiter) SlotOpenBeforeTimeout() bool {
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

//Start waits timeout seconds for a for a slot to open. If a slot opens, it adds 1 to the current requests and returns true.
func (psl *PerSecondLimiter) Start() bool {
	if psl.SlotOpenBeforeTimeout() {
		psl.add()
		return true
	}
	return false
}

//Done signifies that a request is completed; this usually opens up a slot for another.
func (psl *PerSecondLimiter) Done() {
	atomic.AddInt64(&psl.current, -1)
	atomic.AddInt64(&psl.completed, 1)
}

func (psl *PerSecondLimiter) openSlot() bool {
	if atomic.LoadInt64(&psl.current) == 0 {
		psl.startNewRequestCycle()
		return true
	}
	psl.mutex.Lock()
	elapsed := time.Now().Sub(psl.startOfFirstRequest).Seconds()
	psl.mutex.Unlock()
	requestsPerSecond := float64(atomic.LoadInt64(&psl.current)) / elapsed
	return requestsPerSecond < psl.maxPerSecond
}

func (psl *PerSecondLimiter) startNewRequestCycle() {
	atomic.StoreInt64(&psl.requestsThisPeriod, 0)
	psl.mutex.Lock()
	psl.startOfFirstRequest = time.Now()
	psl.mutex.Unlock()
}

//Wait waits for all current requests to complete. This can wait forever if some requests never finish.
//Analagous to sync.WaitGroup.Wait()
//Use WaitTimeOut if that's not what you want.
func (psl *PerSecondLimiter) Wait() {
	for atomic.LoadInt64(&psl.current) > 0 {
		time.Sleep(psl.pollingPeriod)
	}
}

//WaitTimeout returns true if all current requests finish before timeout.
func (psl *PerSecondLimiter) WaitTimeout(timeout time.Duration) bool {
	for wait := time.Duration(0); wait < psl.timeout; wait += psl.pollingPeriod {
		if atomic.LoadInt64(&psl.current) == 0 {
			return true
		}
		time.Sleep(psl.pollingPeriod)
	}
	return false
}
