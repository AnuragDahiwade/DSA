package tree

// Leetcode 863
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parent := make(map[*TreeNode]*TreeNode)

	var buildParent func(*TreeNode)
	buildParent = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			parent[node.Left] = node
		}

		if node.Right != nil {
			parent[node.Right] = node
		}

		buildParent(node.Left)
		buildParent(node.Right)
	}

	buildParent(root)

	// DFS from target
	visited := make(map[*TreeNode]bool)
	queue := &Queue[*TreeNode]{}
	queue.Enqueue(target)
	visited[target] = true
	dist := 0

	for queue.Size() > 0 {
		if dist == k {
			break
		}
		levelSize := queue.Size()

		for i := 0; i < levelSize; i++ {
			node, _ := queue.Dequeue()

			if node.Left != nil && !visited[node.Left] {
				queue.Enqueue(node.Left)
				visited[node.Left] = true
			}

			if node.Right != nil && !visited[node.Right] {
				queue.Enqueue(node.Right)
				visited[node.Right] = true
			}

			if p, ok := parent[node]; ok && !visited[p] {
				queue.Enqueue(p)
				visited[p] = true
			}
		}
		dist++
	}

	// add result
	res := []int{}

	for !queue.IsEmpty() {
		node, _ := queue.Dequeue()
		res = append(res, node.Val)
	}

	return res
}
