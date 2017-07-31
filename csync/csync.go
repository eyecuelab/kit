package csync

import "time"

var (
	concurrentRequests Queue
	Start              Signal
)

const (
	maxRequests = 500
)

type Queue chan Signal
type Signal struct{}

type SignalFunc func() bool

func init() {
	concurrentRequests = make(Queue, maxRequests)
}

//StartTimeoutRequest waits for a slot to open up in concurrentRequests within timeout seconds. Then, it sends a signal that it is one of the concurrent requests.
//startRequest returns a function that frees up a slot in concurrentRequests. use startRequest by deferring it's return value;
//i.e, defer StartRequest()
//ex:
// given a func doWork that we want to launch in a goroutine, we start doWork as follows:
// go doWork() {defer StartRequest() //function logic here}
func StartTimeoutRequest(timeout time.Duration) SignalFunc {
	if WaitForConcurrentRequests(concurrentRequests, timeout) {
		concurrentRequests <- Start
		return func() bool {
			<-concurrentRequests
			return true
		}
	}
	return FailedRequest
}

func FailedRequest() bool {
	return false
}

func StartForeverRequest() func() {
	concurrentRequests <- Start
	return func() {
		<-concurrentRequests
	}
}

//WaitForConcurrentRequests waits for a slot to open up in the default concurentRequests channel within timeout duration.
//It returns true if a slot opens up in time and false otherwise.
func WaitForConcurrentRequests(requests Queue, timeout time.Duration) bool {
	var t time.Duration
	for ; len(requests) == cap(requests); t += 100 * time.Millisecond {
		time.Sleep(100 * time.Millisecond)
		if t > timeout {
			return false
		}
	}
	return true
}

func WaitForeverForConcurrentRequests(requests Queue) {
	for len(requests) == cap(requests) {
		time.Sleep(100 * time.Millisecond)
	}
}

func NewChannelWithTimeout(max int, timeout time.Duration) (maxRequests Queue, startRequest SignalFunc) {
	concurrentRequests := make(Queue, max)
	if WaitForConcurrentRequests(concurrentRequests, timeout) {
		concurrentRequests <- Start
		return concurrentRequests, func() bool {
			<-concurrentRequests
			return true
		}
	}
	return concurrentRequests, FailedRequest
}
