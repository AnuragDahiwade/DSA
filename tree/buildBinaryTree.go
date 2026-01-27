package tree

// Build from the preOrder sequence ********
func buildBinaryTree(arr []int) *TreeNode {
	root := &TreeNode{}
	if len(arr) == 0 {
		return root
	}
	root, _ = helperBuild(&arr, 0)
	return root
}
func helperBuild(arr *[]int, idx int) (*TreeNode, int) {
	if (*arr)[idx] == -1 {
		return nil, idx
	}

	node := &TreeNode{Val: (*arr)[idx], Left: nil, Right: nil}
	node.Left, idx = helperBuild(arr, idx+1)
	node.Right, idx = helperBuild(arr, idx+1)

	return node, idx
}
