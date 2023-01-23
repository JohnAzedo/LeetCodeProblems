package easy

import "github.com/johnazedo/interview-questions/leetcode"

/*
Difficult: Easy
Link: https://leetcode.com/problems/reverse-linked-list
*/
func reverseList(head *leetcode.ListNode) *leetcode.ListNode {
	// Time: O(N)
	// Space: O(N)

	// TODO: Do this with O(1) space complexity
	stack := []int{}
	temp := head

	for temp != nil {
		stack = append(stack, temp.Val)
		temp = temp.Next
	}

	temp = head

	for i := len(stack) - 1; i >= 0; i-- {
		temp.Val = stack[i]
		temp = temp.Next
	}

	return head
}