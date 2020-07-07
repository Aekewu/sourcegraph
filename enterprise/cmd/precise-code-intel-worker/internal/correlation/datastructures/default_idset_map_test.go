package datastructures

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDefaultIDSetMap(t *testing.T) {
	m := DefaultIDSetMap{}
	m.GetOrCreate(50).Add(51)
	m.GetOrCreate(50).Add(52)
	m.GetOrCreate(51).Add(53)
	m.GetOrCreate(51).Add(54)

	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	expected := []int{50, 51}
	if diff := cmp.Diff(expected, keys); diff != "" {
		t.Errorf("unexpected keys (-want +got):\n%s", diff)
	}

	expected = []int{51, 52}
	if diff := cmp.Diff(expected, m[50].Values(nil)); diff != "" {
		t.Errorf("unexpected keys (-want +got):\n%s", diff)
	}

	expected = []int{53, 54}
	if diff := cmp.Diff(expected, m[51].Values(nil)); diff != "" {
		t.Errorf("unexpected keys (-want +got):\n%s", diff)
	}
}
