//package limitrequests contains tools for limiting the number of simulatenously launched requests.
package limitrequests

import (
	"sync/atomic"
	"time"
)

func min(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}

//TotalLimiter limits the total number of current requests. Usage pattern is follows:
//suppose we want to run at most 500 of a func networkCall.
//const maxSimulataneousRequests, timeOut = 500, 30*time.Second
//var startRequest, endRequest = TotalRequestLimiter(maxSimulataneousRequests, timeOut)
//for request := range requests {
//	if startRequest() {
//		go func(request) {
//			defer endRequest
//			networkCall(request)
//		}(request)
//	} else {
//		log.Fatal(fmt.Errorf("could not launch networkCall(%v): timeout after %d seconds", maxSimulataneousRequests, timeout.Seconds())
//	}
//
//}
func TotalLimiter(maxRequests int64, timeout time.Duration) (startRequest func() bool, finishRequest func()) {
	var currentRequests *int64
	waitForRequests := func() bool {
		for wait, i := time.Duration(0), uint(0); wait < timeout; i++ {
			if atomic.LoadInt64(currentRequests) < maxRequests {
				return true
			}
			wait += min(10*time.Millisecond<<i, 250*time.Millisecond)
		}
		return false
	}

	startRequest = func() bool {
		if timeout := !waitForRequests(); timeout {
			return false
		}

		atomic.AddInt64(currentRequests, 1)
		return true
	}
	finishRequest = func() {
		atomic.AddInt64(currentRequests, -1)
	}
	return startRequest, finishRequest
}

//PerSecondLimiter is for situations where the total number of current requests is less important than the number of requests launched per second.
//Note that PerSecondLimiter cannot lock.
func PerSecondLimiter(maxRequestsPerSecond float64, pollingPeriod time.Duration) (startRequest func()) {
	if maxRequestsPerSecond <= 0 {
		panic("must enter a positive number for maxRequestsPerSecond")
	}
	if pollingPeriod <= 0 {
		panic("must enter a positive duration for pollingPeriod")
	}
	startTime := time.Now()
	var totalRequests *uint64
	return func() {
		for {
			elapsed := time.Now().Sub(startTime).Seconds()
			requestsPerSecond := float64(atomic.LoadUint64(totalRequests)) / elapsed
			if requestsPerSecond < maxRequestsPerSecond {
				atomic.AddUint64(totalRequests, 1)
				return
			}
			time.Sleep(pollingPeriod)
		}
	}
}
