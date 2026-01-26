package main

/*
	stack := NewStack[int]
*/

type Stack[T any] struct {
	nums []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) push(num T) {
	s.nums = append(s.nums, num)
}

func (s *Stack[T]) pop() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	zero = s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return zero, true
}

func (s *Stack[T]) peek() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	zero = s.nums[len(s.nums)-1]
	return zero, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.nums) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.nums)
}

// ************************************************************************************

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
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
