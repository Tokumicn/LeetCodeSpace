package main

import (
	"LeetCodeSpace/stack"
	"fmt"
)

// 用两个栈实现队列

type CQueue struct {
	in, out stack.Stack
}

func Constructor() CQueue {
	var q CQueue
	return q
}

func (this *CQueue) AppendTail(value int) {
	this.in.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.out == nil {
		if this.in == nil {
			return -1
		}

		for this.in != nil {
			this.out.Push(this.in.Pop())
		}
		return this.out.Pop().(int)
	} else {
		return this.out.Pop().(int)
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	q := Constructor()
	q.AppendTail(3)
	fmt.Println(q.DeleteHead())

}
