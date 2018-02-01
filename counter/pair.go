package counter

//Pair represents a Key, Count pair
type Pair struct {
	Key   string
	Count int
}

//Pairs is a slice of Pairs. It is sortable lexigraphically by count, then key.
type Pairs []Pair

func (p Pair) KeyVal() (string, int) { return p.Key, p.Count }

func (pairs Pairs) Less(i, j int) bool {
	a, b := pairs[i].Count, pairs[j].Count
	return a < b || (a <= b && pairs[i].Key < pairs[j].Key)
}
func (pairs Pairs) Len() int      { return len(pairs) }
func (pairs Pairs) Swap(i, j int) { pairs[i], pairs[j] = pairs[j], pairs[i] }

func (pairs Pairs) Keys() []string {
	keys := make([]string, len(pairs))
	for i, p := range pairs {
		keys[i] = p.Key
	}
	return keys
}
func (pairs Pairs) Counts() []int {
	counts := make([]int, len(pairs))
	for i, c := range pairs {
		counts[i] = c.Count
	}
	return counts
}
