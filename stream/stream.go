package stream

import (
	"sync"
)

type signal struct{}

var yes signal

func mergeUniqueStrings(chans ...<-chan string) chan string {
	uniques := make(chan string)
	wg := &sync.WaitGroup{}
	for _, ch := range chans {
		wg.Add(1)
		seen := make(map[string]signal)
		go func(ch <-chan string) {
			defer wg.Done()
			for str := range ch {
				if _, ok := seen[str]; !ok {
					seen[str] = yes
					uniques <- str
				}
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		defer close(uniques)
	}()

	return uniques
}
