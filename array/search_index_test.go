package array

import "testing"

func TestSearchIndex(t *testing.T) {
	s := []int{1, 3, 5, 6}
	want := 2
	got := searchInsert(s, 5)
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
