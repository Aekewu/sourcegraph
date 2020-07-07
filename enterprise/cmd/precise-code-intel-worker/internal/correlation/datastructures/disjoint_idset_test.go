package datastructures

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDisjointIDSet(t *testing.T) {
	s := DisjointIDSet{}
	s.Union(1, 2)
	s.Union(3, 4)
	s.Union(1, 3)
	s.Union(5, 6)

	if diff := cmp.Diff([]int{1, 2, 3, 4}, s.ExtractSet(1).Values(nil)); diff != "" {
		t.Errorf("unexpected keys (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff([]int{1, 2, 3, 4}, s.ExtractSet(2).Values(nil)); diff != "" {
		t.Errorf("unexpected keys (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff([]int{1, 2, 3, 4}, s.ExtractSet(3).Values(nil)); diff != "" {
		t.Errorf("unexpected keys (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff([]int{1, 2, 3, 4}, s.ExtractSet(4).Values(nil)); diff != "" {
		t.Errorf("unexpected keys (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff([]int{5, 6}, s.ExtractSet(5).Values(nil)); diff != "" {
		t.Errorf("unexpected keys (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff([]int{5, 6}, s.ExtractSet(6).Values(nil)); diff != "" {
		t.Errorf("unexpected keys (-want +got):\n%s", diff)
	}
}
