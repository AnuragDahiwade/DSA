package tree

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	curr := root

	for curr != nil {
		if curr.Left != nil {
			pred := curr.Left

			for pred.Right != nil && pred.Right != curr {
				pred = pred.Right
			}

			if pred.Right == nil {
				res = append(res, curr.Val)
				pred.Right = curr
				curr = curr.Left
			} else {
				pred.Right = nil
				curr = curr.Right
			}

		} else {
			res = append(res, curr.Val)
			curr = curr.Right
		}
	}

	return res
}
