package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(missingNumber([]int{0}))
	//fmt.Println(missingNumber([]int{1}))
	//fmt.Println(missingNumber([]int{1,2,3}))
	//fmt.Println(missingNumber([]int{0,1,3}))

	//fmt.Println(findIndex([]int{5,7,7,8,8,10}, 8))
	//
	//fmt.Println(findIndex([]int{8,8,10}, 8))
	//
	//fmt.Println(findIndex([]int{8,8,10}, 11))

	//fmt.Println(findIndex([]int{8,8,10}, 9))
	//fmt.Println(sort.SearchInts([]int{8,8,10}, 8+1))

	//fmt.Println(search([]int{1}, 1))
	//fmt.Println(search2([]int{1}, 1))

	//fmt.Println(findIndex([]int{1}, 1))
	//fmt.Println(findIndex([]int{1}, 2))
	//fmt.Println(sort.SearchInts([]int{1}, 2))

	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
}

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1

	for l < r {
		mid := ((r - l) + l) / 2

		if numbers[mid] < numbers[r] { // 右侧有序，最小值在左侧
			r = mid
		} else if numbers[mid] > numbers[r] { // 左侧有序，最小值在右侧
			l = mid + 1
		} else { // 解决全部相等的情况
			r--
		}
	}
	return numbers[l]
}

func search2(nums []int, target int) int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return rightmost - leftmost + 1
}

func search(nums []int, target int) int {

	firstIndex := findIndex(nums, target)

	if firstIndex == len(nums) || nums[firstIndex] != target {
		return 0
	}

	lastIndex := findIndex(nums, target+1) - 1

	return lastIndex - firstIndex + 1
}

func findIndex(nums []int, target int) int {

	i, j := 0, len(nums)
	for i < j {
		m := (i + j) / 2 //int(uint(i+j) >> 1)

		if nums[m] >= target {
			j = m
		} else {
			i = m + 1
		}
	}

	return i
}

func findIndex2(nums []int, target int) int {

	i, j := 0, len(nums)-1
	for i < j {
		m := (i + j) / 2

		if nums[i] >= target {
			j = m
		} else {
			i = m + 1
		}
	}

	return i
}

func missingNumber(nums []int) int {
	var i, j = 0, len(nums) - 1

	for i <= j {
		m := (i + j) / 2

		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	return i
}

func missingNumber2(nums []int) int {

	// nil : -1
	if len(nums) <= 0 {
		return -1
	}
	// [0] : 1
	if len(nums) == 0 && nums[0] == 0 {
		return 1
	}
	// [1] : 0
	if len(nums) == 0 && nums[0] == 1 {
		return 0
	}

	// 过程缺失
	// 过程不缺 -- 前缺   res = nums[0] - 1
	// 过程不缺 -- 后缺   res = nums[len(nums)-1] + 1
	return IsContinuous(nums)
}

func IsContinuous(nums []int) int {

	prev := nums[0]
	length := len(nums)
	for i := 1; i < length; i++ {
		if nums[i]-prev == 1 { // 前后相差1
			prev = nums[i]
			continue
		} else {
			// 过程缺失
			return prev + 1
		}
	}

	// 过程不缺 -- 前缺   res = nums[0] - 1
	if nums[0] > 0 {
		return nums[0] - 1
	}
	// 过程不缺 -- 后缺   res = nums[len(nums)-1] + 1
	if nums[0] == 0 {
		return nums[length] + 1
	}

	return -1
}
