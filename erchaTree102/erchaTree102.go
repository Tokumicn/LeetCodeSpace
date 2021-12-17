package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	s := "00000-42"
	s = strings.Trim(s, "0")
	atoi, _ := strconv.Atoi(s)
	fmt.Println(atoi)
}

func myAtoi(s string) int {

	buf := strings.Builder{}
	sign := 1

	for i := 0; i < len(s); i++ {
		// 前置空格去掉
		if s[i] == ' ' {
			continue
		}

		// 正负号确认
		if s[i] == '-' {
			sign = -1
			continue
		}

		if s[i] >= '0' && s[i] <= '9' {
			buf.WriteByte(s[i])
		} else {
			break
		}
	}

	atoi, err := strconv.Atoi(buf.String())
	if err != nil {
		return 0
	}

	if atoi > math.MaxInt32 {
		atoi = math.MaxInt32
	}

	return atoi * sign
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

var res [][]int

func levelOrder2(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root != nil {
		if len(res) == level {
			res = append(res, []int{})
		}

		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
}
