package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	// --  @# addTwoNumbers
	/*l1 := &ListNode{2, &ListNode{4, nil}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	r := addTwoNumbers(l1, l2)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}*/

	// --  @# lengthForLongestSubstring
	/*str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))*/
	//a := append([]int{}, 1)
	//fmt.Println(min(a[0]))
	/*s := []string{"1111", "2222"}

	count := 0
	for range iter0(s) {
		count++
	}
	fmt.Println(count)

	for x := range slices.Values(s) {
		fmt.Println(x)
	}

	for x, y := range slices.Backward(s) {
		fmt.Println(x, y)
	}*/

	s := []int{1, 1, 1, 1, 4, 6, 4, 6, 7, 7, 7, 7, 9}
	sort.Ints(s)
	r := slices.Compact(s)
	fmt.Println(r, len(r), cap(r))
	r2 := slices.Grow(r, 10)
	fmt.Println(r2, len(r2), cap(r2))

}
