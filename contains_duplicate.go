package main

func main() {

}

// 存在重复

// 给定一个整数数组，判断是否存在重复元素。
// 如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
func containsDuplicate(nums []int) bool {
	if nums == nil || len(nums) == 1 {
		return false
	}

	length := len(nums)
	maps := make(map[int]int, length)
	for _, value := range nums {
		if maps[value] == 0 {
			maps[value] += 1
		} else {
			return true
		}
	}
	return false
}
