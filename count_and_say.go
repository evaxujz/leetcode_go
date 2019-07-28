package main

import "strconv"

func main() {

}

// 报数

// 报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
//
// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221
// 1 被读作  "one 1"  ("一个一") , 即 11。
// 11 被读作 "two 1s" ("两个一"）, 即 21。
// 21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
//
// 给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
//
// 注意：整数顺序将表示为一个字符串。
func countAndSay(n int) string {
	nums := []int{1}
	if n == 1 {
		return "1"
	}

	for i := 2; i <= n; i++ {
		nextNums := make([]int, 0)
		flag := -1
		count := 0
		for _, num := range nums {
			if flag == -1 {
				flag = num
				count++
				continue
			}

			if flag == num {
				count++
			} else {
				nextNums = append(nextNums, count)
				nextNums = append(nextNums, flag)
				flag = num
				count = 1
			}
		}
		nextNums = append(nextNums, count)
		nextNums = append(nextNums, flag)
		nums = nextNums
	}
	say := ""
	for _, value := range nums {
		say += strconv.Itoa(value)
	}
	return say
}
