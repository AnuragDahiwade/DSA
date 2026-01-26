package main

type Queueue[T any] struct {
	data []T
}

func New[T any]() *Queueue[T] {
	return &Queueue[T]{}
}

func (q *Queueue[T]) Enqueue(val T) {
	q.data = append(q.data, val)
}

func (q *Queueue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.data) == 0 {
		return zero, false
	}

	val := q.data[0]
	q.data = q.data[1:]
	return val, true
}

func (q *Queueue[T]) Peek() (T, bool) {
	var zero T
	if len(q.data) == 0 {
		return zero, false
	}
	return q.data[0], true
}

func (q *Queueue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queueue[T]) Size() int {
	return len(q.data)
}
