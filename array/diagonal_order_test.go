package array

import (
	"reflect"
	"testing"
)

func TestDiagonalOrder(t *testing.T) {
	ints := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	got := findDiagonalOrder(ints)
	want := []int{1, 2, 4, 7, 5, 3, 6, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}
