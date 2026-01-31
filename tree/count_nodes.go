package tree

func countNodes(root *TreeNode) int {
	return DFS(root) - 1
}

func DFS(root *TreeNode) int {
	if root == nil {
		return 1
	}

	return DFS(root.Left) + DFS(root.Right)
}
