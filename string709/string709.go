package main

import (
	"fmt"
	"strings"
)

func main() {
	var f float32 = 7.1
	fmt.Printf("%v", f)

	//fmt.Println('a'-'A')
	s := "Hello how are you Contestant"
	//s = strings.Trim(s, " ")
	//fmt.Println(toLower2(s))

	//strings.ToLower(s)

	n := strings.Index(s, " ")
	fmt.Println(n)

}

//func toLowerCase(s string) string {
//	//return strings.ToLower(s)
//
//}

func toLower(s string) string {
	res := make([]byte, len(s))
	for _, ch := range s {
		if ch < 'a' { // 是否为大写
			res = append(res, byte(ch+32))
		} else {
			res = append(res, byte(ch))
		}
	}

	return string(res)
}

func toLower2(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for i, _ := range s {
		ch := s[i]
		if 'A' <= ch && ch <= 'Z' {
			ch += 'a' - 'A'
		}
		b.WriteByte(ch)
	}

	return b.String()
}
