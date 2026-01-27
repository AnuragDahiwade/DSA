package tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := New[*TreeNode]()
	res := [][]int{}

	queue.Enqueue(root)
	for !queue.IsEmpty() {

		levelSize := queue.Size()
		level := []int{}

		for i := 0; i < levelSize; i++ {
			if node, ok := queue.Dequeue(); ok {
				level = append(level, node.Val)

				if node.Left != nil {
					queue.Enqueue(node.Left)
				}
				if node.Right != nil {
					queue.Enqueue(node.Right)
				}
			}

		}
		res = append(res, level)
	}

	return res
}

type Queue[T any] struct {
	data []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(val T) {
	q.data = append(q.data, val)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.data) == 0 {
		return zero, false
	}

	val := q.data[0]
	q.data = q.data[1:]
	return val, true
}

func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if len(q.data) == 0 {
		return zero, false
	}
	return q.data[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}
