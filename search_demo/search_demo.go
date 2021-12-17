package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right - left) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	return -1
}
