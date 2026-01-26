package main

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

func (this *MyStack) Push(x int) {
	secQueue := []int{}

	secQueue = append(secQueue, x)

	for i := 0; i < len(this.queue); i++ {
		secQueue = append(secQueue, this.queue[i])
	}

	this.queue = []int{}
	for i := 0; i < len(secQueue); i++ {
		this.queue = append(this.queue, secQueue[i])
	}
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
