package main

import (
	"fmt"
)

type Stacck[T any] struct {
	nums []T
}

func (s *Stacck[T]) push(num T) {
	s.nums = append(s.nums, num)
}

func (s *Stacck[T]) pop() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	zero = s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return zero, true
}

func (s *Stacck[T]) peek() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	zero = s.nums[len(s.nums)-1]
	return zero, true
}

func (s *Stacck[T]) IsEmpty() bool {
	return len(s.nums) == 0
}

func genMain() {
	stack := Stacck[int]{}
	stack.push(11)
	stack.push(22)
	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
