package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(strStr("mississippi","issip"))

	fmt.Println(strStr("mississippi", "issipi"))

}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i, _ := range haystack {
		if len(haystack)-i < len(needle) {
			return -1
		}

		if haystack[i] == needle[0] {
			if subCheck(haystack[i:], needle) {
				return i
			}
		}
	}
	return -1
}

func subCheck(subStr string, needle string) bool {
	tempSub := subStr[:len(needle)]
	return tempSub == needle
}

func isPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if !checkRune(s[left]) {
			left++
			continue
		}

		if !checkRune(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func checkRune(c byte) bool {
	if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
		return true
	}
	return false
}
