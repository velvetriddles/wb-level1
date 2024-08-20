package main

import (
	"fmt"
)

func isUniq(str string) bool {
	// str = strings.ToLower(str)
	mp := make(map[rune]struct{}, len(str))
	for _, v := range str {
		_, ok := mp[v]
		if ok {
			return false
		}
		mp[v] = struct{}{}
	}
	return true
}

func main() {
	str1 := "abcd"
	fmt.Println(str1, isUniq(str1))
	str2 := "fAaF"
	fmt.Println(str2, isUniq(str2))
	str3 := "aabcd"
	fmt.Println(str3, isUniq(str3))
}
