package csync

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	concurrentRequests Queue
	Start              Signal
)

const (
	maxRequests = 500
)

type Queue chan Signal

func (q *Queue) Full() bool {
	return cap(*q) == len(*q)
}

type Signal struct{}

func init() {
	concurrentRequests = make(Queue, maxRequests)
}

//WaitForConcurrentRequests waits for a slot to open up in the default concurentRequests channel within timeout duration.
//It returns true if a slot opens up in time and false otherwise.
func WaitForConcurrentRequests(requests Queue, timeout time.Duration) bool {
	for t := time.Duration(0); requests.Full(); t += time.Millisecond * time.Duration(rand.Intn(1000)) {
		if t > timeout {
			return false
		}
		time.Sleep(t)
	}
	return true
}

func WaitForeverForConcurrentRequests(requests Queue) {
	for requests.Full() {
		time.Sleep(100 * time.Millisecond)
	}
}

//NO_OP does nothing.
func NO_OP() {}

func NewChannelWithTimeout(max int, timeout time.Duration, errors chan error) (concurrentRequests Queue, startRequest func(caller string) func()) {
	concurrentRequests = make(Queue, max)
	startRequest = func(caller string) func() {
		if WaitForConcurrentRequests(concurrentRequests, timeout) {
			concurrentRequests <- Start
			return func() {
				<-concurrentRequests
			}
		}
		errors <- fmt.Errorf("%s: NewChannelWithTimeout: timeout after %f seconds", caller, timeout.Seconds())
		return NO_OP
	}
	return concurrentRequests, startRequest
}
