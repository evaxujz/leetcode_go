package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("abscfsefsa"))
}

// 字符串中的第一个唯一字符

// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}

	letters := make(map[rune]int, 26)
	for index, value := range s {
		if letters[value] != 0 {
			letters[value] = -1
		} else {
			letters[value] = index + 1
		}
	}

	index := -1
	for _, value := range letters {
		if value != -1 {
			if index == -1 {
				index = value
			} else if value < index {
				index = value
			}
		}
	}

	if index != -1 {
		return index - 1
	} else {
		return index
	}
}
