package main

import "fmt"

func main() {
	list := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	fmtList := list
	for fmtList != nil {
		fmt.Printf("%d -> ", fmtList.Val)
		fmtList = fmtList.Next
	}

	fmt.Println()
	newList := reverseList(list)
	for newList != nil {
		fmt.Printf("%d -> ", newList.Val)
		newList = newList.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
