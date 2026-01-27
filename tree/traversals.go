package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ItarativePreOrder(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}
	stack := NewStack[*TreeNode]()
	stack.Push(root)

	for !stack.IsEmpty() {
		node, _ := stack.Peek()
		res = append(res, node.Val)

		stack.Pop()

		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
	return res
}

func IterativeInorder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	node := root
	st := NewStack[*TreeNode]()
	for true {
		if node != nil {
			st.Push(node)
			node = node.Left
		} else {
			if st.IsEmpty() {
				break
			}

			node, _ := st.Peek()
			res = append(res, node.Val)
			node = node.Right
		}
	}
	return res
}

type Stack[T any] struct {
	nums []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	s.nums = append(s.nums, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	val := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return val, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	return s.nums[len(s.nums)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.nums) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.nums)
}
