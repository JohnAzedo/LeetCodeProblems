package easy

import (
	. "github.com/johnazedo/interview-questions/leetcode"
)

/*
Number: 94
Difficult: Easy
Link: https://leetcode.com/problems/binary-tree-inorder-traversal/
Tags: Stack, Tree, Depth-First Search
Status: Reviewed
*/
func inorderTraversal(root *TreeNode) []int {
	// N is the number of nodes and H is the height of the tree
	// Time: O(N)
	// Space: O(H)
	var result []int
	recursive(root, &result)
	return result
}

func recursive(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	recursive(root.Left, result)
	*result = append(*result, root.Val)
	recursive(root.Right, result)
}

func iterative(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			n := len(stack) - 1
			root = stack[n]
			stack = stack[:n]
			result = append(result, root.Val)
			root = root.Right
		}
	}
	return result
}