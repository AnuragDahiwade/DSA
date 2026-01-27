package main

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	ppath := []*TreeNode{}
	qpath := []*TreeNode{}
	path := []*TreeNode{}

	rootLR_new(root, p, q, &path, &ppath, &qpath)

	i, j := 0, 0
	for i < len(ppath) && j < len(qpath) && ppath[i] == qpath[j] {
		i++
		j++
	}

	return ppath[i-1]
}

func rootLR_new(root, p, q *TreeNode, path, ppath, qpath *[]*TreeNode) {
	if root == nil {
		return
	}

	*path = append(*path, root)

	if root == p {
		*ppath = append([]*TreeNode{}, (*path)...)
	}
	if root == q {
		*qpath = append([]*TreeNode{}, (*path)...)
	}

	// Continue DFS
	rootLR_new(root.Left, p, q, path, ppath, qpath)
	rootLR_new(root.Right, p, q, path, ppath, qpath)

	// Backtrack (remove last)
	*path = (*path)[:len(*path)-1]
}
