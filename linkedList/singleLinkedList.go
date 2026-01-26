package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type SingleLL struct {
	Head *Node
}

func NewNode(num int) *Node {
	return &Node{Data: num, Next: nil}
}

func (s *SingleLL) AddNode(num int) {
	node := NewNode(num)

	if s.Head == nil {
		s.Head = node
	} else {
		temp := s.Head
		for temp.Next != nil {
			temp = temp.Next
		}

		temp.Next = node
	}
}

func (s *SingleLL) RemoveNode(num int) {
	if s.Head == nil {
		fmt.Println("List is empty. so cant remove the node....")
		return
	}

	if s.Head.Data == num {
		s.Head = nil
		return
	}

	// Traverse the list
	prev := s.Head
	temp := s.Head.Next

	for temp != nil && temp.Data != num {
		prev = temp
		temp = temp.Next
	}

	if temp == nil {
		fmt.Println("Node not found in the list.")
		return
	}

	prev.Next = temp.Next

}

func (s *SingleLL) printLL() {
	temp := s.Head
	for temp != nil {
		fmt.Printf("data: %d --> ", temp.Data)
		temp = temp.Next
	}
	fmt.Println()
}

func implementSingleLL() {
	singelLL := SingleLL{}
	singelLL.AddNode(10)
	singelLL.AddNode(20)
	singelLL.AddNode(30)
	singelLL.printLL()
	singelLL.RemoveNode(20)
	singelLL.printLL()
}
