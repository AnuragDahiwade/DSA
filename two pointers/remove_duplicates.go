/**
 * Definition for singly-linked list.
 */

//  83. Remove Duplicates from Sorted List

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	front, back := head, head

	for front != nil {
		for front != nil && front.Val == back.Val {
			front = front.Next
		}
		back.Next = front
		back = front
		if front != nil {
			front = front.Next
		}
	}

	return head
}
