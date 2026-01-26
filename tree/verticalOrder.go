/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import "slices"

func verticalTraversal(root *TreeNode) [][]int {
	nodeMap := make(map[int][]int)

	helper(root, 0, 0, &nodeMap)

	keys := []int{}
	for key, _ := range nodeMap {
		keys = append(keys, key)
	}

	res := [][]int{}
	if len(keys) == 0 {
		return res
	}
	slices.Sort(keys)

	for _, key := range keys {
		vals, _ := nodeMap[key]
		// slices.Sort(vals)
		res = append(res, vals)
	}

	return res
}

func helper(root *TreeNode, row, col int, nodeMap *map[int][]int) {
	if root == nil {
		return
	}

	(*nodeMap)[col] = append((*nodeMap)[col], root.Val)

	helper(root.Left, row+1, col-1, nodeMap)
	helper(root.Right, row+1, col+1, nodeMap)

}
