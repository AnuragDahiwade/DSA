package tree

// 662. Maximum Width of Binary Tree
func widthOfBinaryTree(root *TreeNode) int {
	res := 0

	if root == nil {
		return res
	}

	// Queue for nodes
	queue := New[*TreeNode]()

	// Queue for positions
	posQueue := New[int]()

	queue.Enqueue(root)
	posQueue.Enqueue(0)

	for !queue.IsEmpty() {

		levelSize := queue.Size()

		firstPos := 0
		lastPos := 0

		for i := 0; i < levelSize; i++ {

			node, _ := queue.Dequeue()
			pos, _ := posQueue.Dequeue()

			// Normalize to avoid overflow
			if i == 0 {
				firstPos = pos
			}
			lastPos = pos

			pos = pos - firstPos

			if node.Left != nil {
				queue.Enqueue(node.Left)
				posQueue.Enqueue(2 * pos)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
				posQueue.Enqueue(2*pos + 1)
			}
		}

		width := lastPos - firstPos + 1

		if res < width {
			res = width
		}
	}

	return res
}

// type Queue[T any] struct {
// 	data []T
// }
//
// func NewQueue[T any]() *Queue[T] {
// 	return &Queue[T]{}
// }
//
// func (q *Queue[T]) Enqueue(val T) {
// 	q.data = append(q.data, val)
// }
//
// func (q *Queue[T]) Dequeue() (T, bool) {
// 	var zero T
// 	if len(q.data) == 0 {
// 		return zero, false
// 	}
//
// 	val := q.data[0]
// 	q.data = q.data[1:]
// 	return val, true
// }
//
// func (q *Queue[T]) Peek() (T, bool) {
// 	var zero T
// 	if len(q.data) == 0 {
// 		return zero, false
// 	}
// 	return q.data[0], true
// }
//
// func (q *Queue[T]) IsEmpty() bool {
// 	return len(q.data) == 0
// }
//
// func (q *Queue[T]) Size() int {
// 	return len(q.data)
// }
