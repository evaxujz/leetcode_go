package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	element := nums[0]
	anchor := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != element {
			nums[anchor] = nums[i]
			anchor += 1
			element = nums[i]
		}
	}
	return anchor
}
