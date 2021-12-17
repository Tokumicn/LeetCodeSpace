package main

import "fmt"

func main() {
	//nums := []int{1,2,3,4}
	//n := len(nums)
	//res := make([]int, n)

	//for i,num := range nums {
	//	res[n-1-i] = num
	//}
	//
	//fmt.Println(res)

	//for i:=0; i<n; i++{
	//	res[n-1-i] = nums[i]
	//}
	//fmt.Println(res)

	//removeDuplicates([]int{1,1,2})

	//merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)

	fmt.Println(intersect([]int{3, 1, 2}, []int{1, 1}))
}

func intersect(nums1 []int, nums2 []int) []int {

	res := []int{}
	hashMap := map[int]int{}

	if len(nums1) <= len(nums2) {
		for _, n := range nums2 {
			hashMap[n] += 1
		}

		for _, n := range nums1 {
			if (hashMap[n] - 1) >= 0 {
				res = append(res, n)
			}
		}
	} else {
		for _, n := range nums1 {
			hashMap[n] += 1
		}

		for _, n := range nums2 {
			if (hashMap[n] - 1) >= 0 {
				res = append(res, n)
				hashMap[n] -= 1
			}
		}
	}

	return res
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := []int{}

	i, j := 0, 0
	for {
		if i >= m && j >= n {
			break
		}

		if i > m {
			i = m
		}

		if j > n {
			j = n
		}

		if nums1[i] > nums2[j] {
			res = append(res, nums2[j])
			j++
		} else {
			res = append(res, nums1[i])
			i++
		}
	}
	nums1 = res
	fmt.Println(res)
	return
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	left, right := 0, 1

	for right < len(nums) {
		if nums[left] == nums[right] {
			right += 1
			continue
		} else {
			left += 1
			nums[left] = nums[right]
			right += 1
		}
	}

	return left
}
