package main

import "fmt"

func main() {
	fmt.Println(isAnagram("ssaas", "assa"))
}

// 有效的字母异位词

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make(map[rune]int, 26)

	for _, value := range s {
		letters[value] += 1
	}

	for _, value := range t {
		if letters[value] == 0 {
			return false
		} else if letters[value] == 1 {
			delete(letters, value)
		} else {
			letters[value] -= 1
		}
	}

	if len(letters) == 0 {
		return true
	} else {
		return false
	}
}
