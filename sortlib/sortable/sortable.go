package sortable

type (
	Uints   []uint
	Uint64s []uint64
	Int64s  []int64
	Ints    []int
	Bytes   []byte
	Runes   []rune
)

func (a Ints) Less(i, j int) bool { return a[i] < a[j] }
func (a Ints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ints) Len() int           { return len(a) }

func (a Uints) Less(i, j int) bool { return a[i] < a[j] }
func (a Uints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Uints) Len() int           { return len(a) }

func (a Uint64s) Less(i, j int) bool { return a[i] < a[j] }
func (a Uint64s) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Uint64s) Len() int           { return len(a) }

func (a Int64s) Less(i, j int) bool { return a[i] < a[j] }
func (a Int64s) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int64s) Len() int           { return len(a) }

func (b Bytes) Less(i, j int) bool { return b[i] < b[j] }
func (b Bytes) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b Bytes) Len() int           { return len(b) }

func (r Runes) Less(i, j int) bool { return r[i] < r[j] }
func (r Runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Runes) Len() int           { return len(r) }
