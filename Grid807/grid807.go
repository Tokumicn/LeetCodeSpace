package main

// 807. 保持城市天际线 https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline/

import "fmt"

func main() {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	cloMax := make([]int, n)
	rowMax := make([]int, n)

	// 记录每行每列的最大值
	for i, rows := range grid {
		for j, val := range rows {
			rowMax[i] = max(rowMax[i], val)
			cloMax[j] = max(cloMax[j], val)
		}
	}

	res := 0

	// 对于每个元素来说最大可补值为    min(该位置列最大值, 该位置行最大值) - 当前值
	for i, rows := range grid {
		for j, val := range rows {
			res += min(rowMax[i], cloMax[j]) - val
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
