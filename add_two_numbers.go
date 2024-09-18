package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	arg1, arg2 := 0, 0
	var n1 *ListNode = nil
	var n2 *ListNode = nil
	if l1 != nil {
		arg1 = l1.Val
		n1 = l1.Next
	}
	if l2 != nil {
		arg2 = l2.Val
		n2 = l2.Next
	}
	sum = arg1 + arg2
	dev := sum / 10
	extra := sum % 10
	// --  @# 递归结束条件
	if (l1 == nil || l1.Next == nil) && (l2 == nil || l2.Next == nil) {
		// --  @# 递归结束处理进位
		if dev > 0 {
			return &ListNode{extra, &ListNode{1, nil}}
		}
		return &ListNode{extra, nil}
	}
	// --  @# 递归中处理进位
	if nil != n1 {
		n1.Val += dev
	} else {
		n2.Val += dev
	}
	// --  @# 开始递归
	next := addTwoNumbers(n1, n2)
	return &ListNode{extra, next}
}
