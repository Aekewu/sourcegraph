package datastructures

import (
	"sort"

	"github.com/google/go-cmp/cmp"
)

// IDSet is a set of string identifiers.
type IDSet struct {
	m map[int]struct{}
}

func NewIDSet() *IDSet {
	return &IDSet{
		m: map[int]struct{}{},
	}
}

func IDSetWith(values ...int) *IDSet {
	set := NewIDSet()
	for _, value := range values {
		set.Add(value)
	}

	return set
}

// Add adds the given element to the set.
func (set *IDSet) Add(id int) {
	set.m[id] = struct{}{}
}

// AddAll adds the contents of the given set to this set.
func (set *IDSet) AddAll(other *IDSet) {
	if other == nil {
		return
	}

	for k := range other.m {
		set.Add(k)
	}
}

// Contains determines if the given element is a member of this set.
func (set *IDSet) Contains(id int) bool {
	_, ok := set.m[id]
	return ok
}

// Keys returns the sorted contents of this set.
func (set *IDSet) Values(keys []int) []int {
	for k := range set.m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return keys
}

// Choose deterministically returns a key from this set. If the set is
// empty, a false-valued flag is returned.
func (set *IDSet) Choose() (int, bool) {
	// TODO - can do this more efficiently now

	if len(set.m) == 0 {
		return 0, false
	}

	return set.Values(nil)[0], true
}

// TODO - document
func (set *IDSet) TakeMin(val *int) bool {
	min, ok := set.Choose()
	if ok {
		delete(set.m, min)
		*val = min
	}

	return ok
}

var IDSetComparer = cmp.Comparer(func(x, y *IDSet) bool {
	if x == nil {
		return y == nil
	}

	if y == nil {
		return false
	}

	vx := x.Values(nil)
	vy := y.Values(nil)

	if len(vx) != len(vy) {
		return false
	}

	for i, v := range vx {
		if vy[i] != v {
			return false
		}
	}

	return true
})
