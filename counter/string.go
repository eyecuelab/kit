package counter

import (
	"fmt"
	"strings"
)

//String counts the occurence of strings.
type String map[string]int

func (counter String) Add(items ...string) {
	for _, item := range items {
		if _, ok := counter[item]; ok {
			counter[item]++
		} else {
			counter[item] = 1
		}
	}
}

func (counter String) String() string {
	formatted := make([]string, len(counter))
	i := 0
	for key, val := range counter {
		formatted[i] = fmt.Sprintf("%s:%d", key, val)
		i++
	}
	return strings.Join(formatted, ", ") + "\n"

}

func CombineStringCounters(counters ...String) String {
	counter := make(String)
	for _, c := range counters {
		for key, val := range c {
			if _, ok := counter[key]; ok {
				counter[key] += val
			} else {
				counter[key] = val
			}
		}
	}
	return counter
}
func Min(counters ...String) String {
	min := make(String)
	for _, c := range counters {
		for key, val := range c {
			if _, ok := min[key]; ok {
				if val < min[key] {
					min[key] = val
				}
			} else {
				min[key] = 0
			}
		}
	}
	return min
}
func Max(counters ...String) String {
	max := make(String)
	for _, c := range counters {
		for key, val := range c {
			if _, ok := max[key]; ok {
				if val > max[key] {
					max[key] = val
				}
			} else {
				max[key] = 0
			}
		}
	}
	return max
}

func FromStrings(strings ...string) String {
	counter := make(String)
	for _, str := range strings {
		counter.Add(str)
	}
	return counter
}
