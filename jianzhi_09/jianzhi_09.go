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
	//q := Constructor()
	//q.AppendTail(3)
	//fmt.Println(q.DeleteHead())

	//s := "We are happy."
	//s = strings.ReplaceAll(s, " ", "%20")

	s := "abcdefg"
	n := 2

	res := s[n:] + s[0:n]

	fmt.Println(res)

}

func firstUniqChar(s string) byte {
	c := [26]int{}
	for _, ch := range s {
		c[ch-'a'] += 1
	}

	for _, ch := range s {
		if c[ch-'a'] == 1 {
			return byte(ch)
		}
	}

	return ' '
}
