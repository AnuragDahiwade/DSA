package main

import "fmt"

type Queue struct {
	nums []int
}

func (q *Queue) push(num int) {
	q.nums = append(q.nums, num)
}

func (q *Queue) pop() (int, bool) {
	if q.isEmpty() {
		return -1, false
	}

	ele := q.nums[0]
	q.nums = q.nums[1:]
	return ele, true
}

func (q *Queue) peek() (int, bool) {
	if q.isEmpty() {
		return -1, false
	}
	return q.nums[0], true
}

func (q *Queue) isEmpty() bool {
	return len(q.nums) == 0
}

func (q *Queue) print() {
	if q.isEmpty() {
		fmt.Println("Queue is empty!!!!")
		return
	}

	for i := 0; i < len(q.nums); i++ {
		fmt.Printf("%d -> ", q.nums[i])
	}
}

func implementation() {
	queue := Queue{}

	queue.print()
	queue.push(11)
	queue.push(22)
	queue.push(33)
	queue.pop()
	queue.pop()
	queue.pop()
	queue.pop()
	queue.print()
}
