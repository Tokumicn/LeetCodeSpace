package main

import "fmt"

func main() {
	// fmt.Println(rowFindNum([]int{-5}, -5))

	fmt.Println('0')
	fmt.Println('1')
	fmt.Println('9')
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, rows := range matrix {

		// 优化少选出某些行
		if len(rows) <= 0 {
			return false
		}

		if rows[0] > target {
			continue
		}

		if rows[len(rows)-1] < target {
			continue
		}

		if rowFindNum(rows, target) {
			return true
		}
	}

	return false
}

// 行二分查找
func rowFindNum(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		} else if nums[mid] == target {
			return true
		}
	}
	return false
}
