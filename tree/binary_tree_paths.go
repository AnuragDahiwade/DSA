/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// LEETCODE : 257 - binary tree paths
package main

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}

	helper_1(root, "", &res)
	return res
}

func helper_1(root *TreeNode, curr string, res *[]string) {
	if root.Left == nil && root.Right == nil {
		curr += strconv.Itoa(root.Val)
		*res = append(*res, curr)
		return
	}

	curr += strconv.Itoa(root.Val) + "->"

	if root.Left != nil {
		helper_1(root.Left, curr, res)
	}
	if root.Right != nil {
		helper_1(root.Right, curr, res)
	}
}
