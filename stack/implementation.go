package main

import "fmt"

type Stack struct {
	nums []int
}

func (s *Stack) push(num int) {
	s.nums = append(s.nums, num)
}

func (s *Stack) pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}

	top := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return top, true
}

func (s *Stack) peek() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}

	return s.nums[len(s.nums)-1], true
}

func (s *Stack) isEmpty() bool {
	return len(s.nums) == 0
}

func (s *Stack) print() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Stack: ")
	for i := len(s.nums) - 1; i >= 0; i-- {
		fmt.Printf("%d \n", s.nums[i])
	}
}

func doSomeOperations() {
	stack := Stack{}

	fmt.Println(stack.pop())
	stack.push(11)
	fmt.Println(stack.peek())
	stack.push(12)
	stack.push(13)
	fmt.Println(stack.pop())
	fmt.Print()
	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	stack.push(33)
	stack.push(32)
	stack.print()

}
