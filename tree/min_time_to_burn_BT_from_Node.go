package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
	// track parents of
	parent := make(map[*TreeNode]*TreeNode)
	var target *TreeNode

	var buildParent func(*TreeNode)
	buildParent = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val == start {
			target = node
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

	if target == nil {
		return 0
	}

	// BFS from the target

	visited := make(map[*TreeNode]bool)
	queue := &Queue[*TreeNode]{}
	queue.Enqueue(target)
	visited[target] = true

	dist := -1

	for queue.Size() > 0 {
		size := queue.Size()

		for i := 0; i < size; i++ {
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

	return dist
}
