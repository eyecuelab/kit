package limitrequests

//ConcurrencyLimiter is an interface for limiting the number of simulataneous concurrent requests.
type ConcurrencyLimiter interface {
	add()
	openSlot() bool
	Done()
	SlotOpenBeforeTimeOut() bool
	Wait() //equivalent to sync.WaitGroup.Wait()
	Start() bool
	// if cl.SlotOpenBeforeTimeout(){
	//		cl.add()
	// 		return true
	//		}
	// return false
}

type MaxLimiter interface {
	ConcurrencyLimiter
	HitMaxRequests() bool
}
