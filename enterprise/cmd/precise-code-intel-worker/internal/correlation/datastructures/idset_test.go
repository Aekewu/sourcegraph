package datastructures

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIDSet(t *testing.T) {
	ids1 := NewIDSet()
	ids1.Add(50)
	ids1.Add(54)

	ids2 := NewIDSet()
	ids2.Add(51)
	ids2.Add(52)
	ids2.Add(53)
	ids2.AddAll(ids1)

	expected := []int{50, 51, 52, 53, 54}
	if diff := cmp.Diff(expected, ids2.Values(nil)); diff != "" {
		t.Errorf("unexpected keys (-want +got):\n%s", diff)
	}

	if value, ok := ids2.Choose(); !ok {
		t.Errorf("expected a value to be chosen")
	} else if value != 50 {
		t.Errorf("unexpected chosen value. want=%d have=%d", 50, value)
	}

	ids2.Add(49)
	if value, ok := ids2.Choose(); !ok {
		t.Errorf("expected a value to be chosen")
	} else if value != 49 {
		t.Errorf("unexpected chosen value. want=%d have=%d", 49, value)
	}
}

func TestChooseEmptyIDSet(t *testing.T) {
	ids := NewIDSet()
	if _, ok := ids.Choose(); ok {
		t.Errorf("unexpected ok")
	}
}
