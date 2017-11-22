package counter

import (
	"fmt"
	"sort"
	"strings"
)

//String counts the occurence of strings.
type String map[string]int

//Add the the following strings to the counter.
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

//Keys returns the keys of the counter. Order is not guaranteed.
func (counter String) Keys() []string {
	keys := make([]string, len(counter))
	var i int
	for k := range counter {
		keys[i] = k
		i++
	}
	return keys
}

func (counter String) MostCommon() (string, int) {
	mostCommon, max := "", 0
	for k, v := range counter {
		if v > max {
			mostCommon, max = k, v
		}
	}
	return mostCommon, max
}

//MostCommonN returns the N most common keys. In the case of a tie, it prioritizes the lexigraphically lowest,
//so that the keys contained in counter.MostCommonN(n) will be identical for two equivalent counters.
func (counter String) MostCommonN(n int) (keys []string, counts []int, ok bool) {
	if len(counter) < n {
		return nil, nil, false
	}
	keys, counts = make([]string, n), make([]int, n)
	for k, v := range counter {
		for i, c := range counts {
			if v > c || v == c && k < keys[i] {
				copy(counts[i+1:], counts[i:n-1])
				copy(keys[i+1:], keys[i:n-1])
				keys[i], counts[i] = k, v
				break
			}
		}
	}
	return keys, counts, true
}

//Sorted returns the keys and coutns of the strings in the counter. They are sorted by the count, and then by the key.
//i.e, given {"c": 1, "b": 3, "a":3}, returns {1, 1, 3}, {c, a, b}
func (counter String) Sorted() ([]string, []int) {
	keys := counter.Keys()
	sort.Slice(keys, func(i, j int) bool {
		return counter[keys[i]] < counter[keys[j]] ||
			(counter[keys[i]] == counter[keys[j]] && keys[i] < keys[j])
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
