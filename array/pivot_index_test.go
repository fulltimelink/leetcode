package array

import "testing"

func TestPivotIndex(t *testing.T) {
	s := []int{1, 7, 3, 6, 5, 6}
	result := 3
	r := pivotIndex(s)
	if r != result {
		t.Errorf("wrong result: got %d, want %d", r, result)
	}

	ss := []int{2, 1, -1}
	result = 0
	r = pivotIndex(ss)
	if r != result {
		t.Errorf("wrong result: got %d, want %d", r, result)
	}

	s = []int{-1, -1, -1, -1, -1, -1}
	result = -1
	r = pivotIndex(s)
	if r != result {
		t.Errorf("wrong result: got %d, want %d", r, result)
	}
}
