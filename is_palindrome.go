package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

// 验证回文字符串

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
// 空字符串定义为有效的回文串。
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	regex := regexp.MustCompile(`[^a-zA-Z0-9]`)
	s = regex.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	count := len(s) >> 1
	indexSum := len(s) - 1
	for i := 0; i < count; i++ {
		if s[i] != s[indexSum-i] {
			return false
		}
	}
	return true
}
