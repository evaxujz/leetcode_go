package main

import "fmt"

func main() {
	fmt.Println(strStr("123", "1"))
}

// 实现strStr()

// 实现 strStr() 函数。
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if haystack == "" || len(haystack) < len(needle) {
		return -1
	}

	nLen := len(needle)
	for i := 0; i <= len(haystack)-len(needle); i++ {
		isMatch := true
		for index, round := i, 0; round < nLen; index, round = index+1, round+1 {
			if haystack[index] != needle[round] {
				isMatch = false
				break
			}
		}
		if isMatch {
			return i
		}
	}

	return -1
}
