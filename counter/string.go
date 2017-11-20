package counter

import (
	"fmt"
	"sort"
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
	var i int
	for key, val := range counter {
		formatted[i] = fmt.Sprintf("%s:%d", key, val)
		i++
	}
	return "{" + strings.Join(formatted, ", ") + "}"

}

func (counter String) Keys() []string {
	keys := make([]string, len(counter))
	var i int
	for k := range counter {
		keys[i] = k
		i++
	}
	return keys
}

func (counter String) Sorted() ([]string, []int) {
	keys := counter.Keys()
	sort.Slice(keys, func(i, j int) bool {
		return counter[keys[i]] < counter[keys[j]]
	})
	vals := make([]int, len(keys))
	for i, k := range keys {
		vals[i] = counter[k]
	}
	return keys, vals

}

func (counter String) Combine(counters ...String) String {
	combined := counter.Copy()
	for _, c := range counters {
		for key, val := range c {
			if _, ok := counter[key]; ok {
				combined[key] += val
			} else {
				combined[key] = val
			}
		}
	}
	return combined
}

func (counter String) Equal(other String) bool {
	if len(counter) != len(other) {
		return false
	}
	for k, v := range counter {
		if b, ok := other[k]; !ok || b != v {
			return false
		}

	}
	return true
}

//PositiveElements returns a copy of counter containing all the elements where the count of counter
func (counter String) PositiveElements() String {
	copy := make(String)
	for k, v := range counter {
		if v > 0 {
			copy[k] = v
		}
	}
	return copy
}

func (counter String) Copy() String {
	copy := make(String)
	for k, v := range counter {
		copy[k] = v
	}
	return copy
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
				min[key] = val
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
				max[key] = val
			}
		}
	}
	return max
}

func FromStrings(strings ...string) String {
	counter := make(String)
	counter.Add(strings...)
	return counter
}
