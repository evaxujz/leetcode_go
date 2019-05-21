package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3, 4}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
	fmt.Println(plusOne([]int{0}))
}

// 加一

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
func plusOne(digits []int) []int {
	newDigits := make([]int, 1)
	newDigits = append(newDigits, digits...)
	for i := len(newDigits) - 1; i >= 0; i-- {
		if newDigits[i] == 9 {
			newDigits[i] = 0
		} else {
			newDigits[i] += 1
			break
		}
	}

	for i := range newDigits {
		if newDigits[i] != 0 {
			return newDigits[i:]
		}
	}

	return []int{0}
}
