//Package counter implements various counter types, analagous to Python's collections.Counter
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

//Count returns counter[string] if it exists, or 0 otherwise.
func (counter String) Count(k string) int {
	n, _ := counter[k]
	return n
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

//Combine one or more counters, where the count of each element is the sum of the count of each
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

//Equal returns true if two counters are equal. Two counters are equal if they have the same length and share all keys and values.
func (counter String) Equal(other String) bool {
	if len(counter) != len(other) {
		return false
	}
	for k, v := range counter {
		if v == 0 {
			if b, _ := other[k]; b != v {
				return false
			}
		} else if b, ok := other[k]; !ok || b != v {
			return false
		}

	}
	return true
}

//NonZeroElements returns a copy of the counter with all dummy elements(count zero) removed.
func (counter String) NonZeroElements(other String) String {
	copy := make(String)
	for k, v := range counter {
		if v != 0 {
			copy[k] = v
		}
	}
	return copy
}

//PositiveElements returns a copy of counter containing all the elements where the counter[x] > 0
func (counter String) PositiveElements() String {
	copy := make(String)
	for k, v := range counter {
		if v > 0 {
			copy[k] = v
		}
	}
	return copy
}

//Copy the counter
func (counter String) Copy() String {
	copy := make(String)
	for k, v := range counter {
		copy[k] = v
	}
	return copy
}

//Min returns a counter minC where minC[x] = min(max(c[x], 0) for c in counters)
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

//Max returns a counter maxC where maxC[x] = max(c[x] for c in counters)
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

//FromStrings creates a counter of strings from a slice, counting each member once.
func FromStrings(strings ...string) String {
	counter := make(String)
	counter.Add(strings...)
	return counter
}
