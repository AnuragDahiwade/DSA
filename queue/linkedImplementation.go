package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type Queuee struct {
	Head *Node
	Tail *Node
}

func (q *Queuee) enqueue(num int) {
	node := newNode(num)

	if q.isEmpty() {
		q.Head = node
		q.Tail = node
		return
	}

	q.Tail.Next = node
	q.Tail = node
}

func (q *Queuee) dequeue() (int, bool) {
	if q.isEmpty() {
		fmt.Println("Queue is empty, can't delete")
		return -1, false
	}

	val := q.Head.Val
	q.Head = q.Head.Next
	return val, true
}

func (q *Queuee) peek() (int, bool) {
	if q.isEmpty() {
		fmt.Println("Queue is empty, can't peek")
		return -1, false
	}
	return q.Head.Val, true
}

func (q *Queuee) isEmpty() bool {
	return q.Head == nil
}

func (q *Queuee) print() {
	if q.isEmpty() {
		fmt.Println("Queue is empty, can't print")
		return
	}
	temp := q.Head
	for temp != nil {
		fmt.Printf("%d -> ", temp.Val)
		temp = temp.Next
	}
}

func newNode(num int) *Node {
	return &Node{Val: num, Next: nil}
}

func linkedImpl() {
	queue := Queuee{}
	queue.print()
	queue.enqueue(11)
	queue.enqueue(22)
	queue.enqueue(33)
	// queue.dequeue()
	// queue.dequeue()
	queue.dequeue()
	queue.print()
}
