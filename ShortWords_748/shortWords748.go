package main

import (
	"fmt"
	"unicode"
)

// 748. 最短补全词
// 先统计 licensePlate 中每个字母的出现次数（忽略大小写），然后遍历 words 中的每个单词，
// 若 26 个字母在该单词中的出现次数均不小于在 licensePlate 中的出现次数，则该单词是一个补全词。
// 返回最短且最靠前的补全词。

func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
}

func shortestCompletingWord(licensePlate string, words []string) string {
	res := ""

	// 统计licensePlate每个单词出现的次数
	cnt := [26]int{}
	for _, ch := range licensePlate {
		if unicode.IsLetter(ch) { // 这个unicode.IsLetter()判断是否是字母的方法
			cnt[unicode.ToLower(ch)-'a'] += 1 // 次数加1   rune相减获得
		}
	}

	for _, word := range words {

		if IsMatch(cnt, word) {
			// 获取最短
			if res == "" || len(word) < len(res) {
				res = word
			}
		}
	}

	return res
}

func IsMatch(cnt [26]int, word string) bool {
	c := [26]int{}
	for _, ch := range word {
		c[ch-'a'] += 1
	}

	for i := 0; i < 26; i++ {
		if c[i] < cnt[i] { // 当前词未出现必要的字母，提前结束
			return false
		}
	}
	return true
}
