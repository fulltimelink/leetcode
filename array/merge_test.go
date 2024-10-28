package array

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	s := [][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}
	r := merge(s)
	fmt.Println(r)
}
