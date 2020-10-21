package main

import "fmt"

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}

type MyStack struct {
	stack []int
}

func CreateMyStack() *MyStack {
	return &MyStack{
		stack: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	ret := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return ret
}

func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}

func (this *MyStack) Peek() int {
	return this.stack[len(this.stack)-1]
}

type MyQueue struct {
	stack *MyStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack: CreateMyStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	temp := []int{}
	for !this.stack.Empty() {
		temp = append(temp, this.stack.Pop())
	}

	this.stack.Push(x)
	for i := len(temp) - 1; i >= 0; i-- {
		this.stack.Push(temp[i])
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	ret := this.stack.Pop()
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
