package main

type Count struct {
	Node *TreeNode
	Num  int
}

func allTraversals(root *TreeNode) [][]int {
	pre := []int{}
	in := []int{}
	post := []int{}
	if root == nil {
		return [][]int{pre, in, post}
	}
	st := NewStack[*Count]()
	ct := &Count{Node: root, Num: 1}
	st.Push(ct)

	for !st.IsEmpty() {
		node, _ := st.Pop()
		if node.Num == 1 {
			pre = append(pre, node.Node.Val)
			node.Num++
			st.Push(node)
			if node.Node.Left != nil {
				st.Push(&Count{Node: node.Node.Left, Num: 1})
			}
		} else if node.Num == 2 {
			in = append(in, node.Node.Val)
			node.Num++
			st.Push(node)
			if node.Node.Right != nil {
				st.Push(&Count{Node: node.Node.Right, Num: 1})
			}
		} else {
			post = append(post, node.Node.Val)
		}
	}

	return [][]int{pre, in, post}
}

func BuildTreeFromLevelOrder(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		curr := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			curr.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, curr.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			curr.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, curr.Right)
		}
		i++
	}

	return root
}
