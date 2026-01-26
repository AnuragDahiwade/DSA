package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type Stackk struct {
	Top *Node
}

func newNode(num int) *Node {
	return &Node{Val: num, Next: nil}
}

// we will push at head nad pop at head also
func (s *Stackk) push(num int) {
	node := newNode(num)

	if s.isEmpty() {
		s.Top = node
		return
	}
	node.Next = s.Top
	s.Top = node
}

func (s *Stackk) pop() (int, bool) {
	if s.isEmpty() {
		fmt.Println("Stack is Empty, Nothing to pop form stack")
		return -1, false
	}
	val := s.Top.Val
	s.Top = s.Top.Next
	return val, true
}

func (s *Stackk) peek() (int, bool) {
	if s.isEmpty() {
		fmt.Println("Stack is Empty, Nothing to peek form stack")
		return -1, false
	}
	return s.Top.Val, true
}

func (s *Stackk) print() {
	fmt.Println("Stack: ")
	if s.isEmpty() {
		fmt.Println("Stack is empty....")
	}

	temp := s.Top
	for temp != nil {
		fmt.Printf("%d _ \n", temp.Val)
		temp = temp.Next
	}
}

func (s *Stackk) isEmpty() bool {
	return s.Top == nil
}

func stackLinkedImpl() {
	stack := Stackk{}

	fmt.Println(stack.pop())
	stack.push(11)
	fmt.Println(stack.peek())
	stack.push(12)
	stack.push(13)
	fmt.Println(stack.pop())
	stack.print()
	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	stack.push(33)
	stack.push(32)
	stack.print()
}
