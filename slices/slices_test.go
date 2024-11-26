package slices

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"testing"
)

func TestCompact(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 8, 9, 10}
	ints2 := slices.Compact(ints)
	fmt.Println(ints, len(ints), cap(ints))
	fmt.Println(ints2, len(ints2), cap(ints2))
}

func TestCompactFunc(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 8, 9, 10}
	ints2 := slices.CompactFunc(ints, func(i, j int) bool {
		return i == j
	})
	fmt.Println(ints, len(ints), cap(ints))
	fmt.Println(ints2, len(ints2), cap(ints2))
}

func TestCompare(t *testing.T) {
	names := []string{"Bob", "Alice", "Vera"}
	fmt.Println("Equal:", slices.Compare(names, []string{"Bob", "Alice", "Vera"}))
	fmt.Println("V < X:", slices.Compare(names, []string{"Bob", "Alice", "Xera"}))
	fmt.Println("V > C:", slices.Compare(names, []string{"Bob", "Alice", "Cat"}))
	fmt.Println("3 > 2:", slices.Compare(names, []string{"Bob", "Alice"}))
	fmt.Println("Order:", slices.Compare(names, []string{"Alice", "Bob", "Vera", "Data"}))
}

func TestCompareFunc(t *testing.T) {
	numbers := []int{0, 43, 8}
	strings := []string{"0", "0", "8"}
	result := slices.CompareFunc(numbers, strings, func(n int, s string) int {
		sn, err := strconv.Atoi(s)
		if err != nil {
			return 1
		}
		return cmp.Compare(n, sn)
	})
	fmt.Println(result)
}

func TestRepeat(t *testing.T) {
	sto := []int{1, 2, 3}
	st := slices.Repeat(sto, 0)
	fmt.Println(sto)
	fmt.Println(st == nil, len(st), cap(st))
}

func TestReplace(t *testing.T) {
	names := []string{"Alice", "Bob", "Vera", "Zac"}
	names2 := slices.Replace(names, 1, 3, "Bill")
	fmt.Println(names, len(names))
	fmt.Println(names2, len(names2))
}

func TestReverse(t *testing.T) {
	s := []string{"a", "b", "c"}
	slices.Reverse(s)
	fmt.Println(s)
}
