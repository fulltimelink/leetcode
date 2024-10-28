package array

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	is := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	want := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	rightRotate(is)
	if !reflect.DeepEqual(is, want) {
		t.Errorf("TestRotate failed. got: %v, want: %v", is, want)
	}
	is = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	want = [][]int{
		{3, 6, 9},
		{2, 5, 8},
		{1, 4, 7},
	}
	leftRotate(is)
	if !reflect.DeepEqual(is, want) {
		t.Errorf("TestRotate failed. got: %v, want: %v", is, want)
	}
	is = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	want = [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	rotate(is, PrincipalDiagonal)
	if !reflect.DeepEqual(is, want) {
		t.Errorf("PrincipalDiagonal failed. got: %v, want: %v", is, want)
	}
	is = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	want = [][]int{
		{9, 6, 3},
		{8, 5, 2},
		{7, 4, 1},
	}
	rotate(is, AuxiliaryDiagonal)
	if !reflect.DeepEqual(is, want) {
		t.Errorf("AuxiliaryDiagonal failed. got: %v, want: %v", is, want)
	}

}
