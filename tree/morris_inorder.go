package tree

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	current_root := root

	for current_root != nil {
		if current_root.Left != nil {
			right := current_root.Left

			for right.Right != nil {
				right = right.Right
			}

			right.Right = current_root

			temp := current_root
			current_root = current_root.Left

			temp.Left = nil
		} else {
			res = append(res, current_root.Val)
			current_root = current_root.Right
		}
	}

	return res
}
