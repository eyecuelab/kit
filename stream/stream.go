package stream

import (
	"sync"
)

type signal struct{}

var yes signal

//MergeUniqueStrings returns a channel which streams the unique output of
//all of the strings in the input channel. The order of output is not guaranteed.
func MergeUniqueStrings(chans ...chan string) chan string {
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

//MergeStrings returns a channel which merges the input of all of the strings in it's input.
//The order of output is not guaranteed.
func MergeStrings(chans ...<-chan string) chan string {
	out := make(chan string)
	wg := &sync.WaitGroup{}
	for _, ch := range chans {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for str := range ch {
				out <- str
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		defer close(out)
	}()
	return out
}
