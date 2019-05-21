package main

import "fmt"

func main() {

}

// 两个数组的交集 II

// 给定两个数组，编写一个函数来计算它们的交集。
func intersect(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int, len(nums1))
	for _, value := range nums1 {
		map1[value] += 1
	}
	nums3 := make([]int, 0)
	fmt.Println(nums3)
	for _, value := range nums2 {
		if map1[value] != 0 {
			nums3 = append(nums3, value)
			fmt.Println(nums3)
			map1[value] -= 1
		}
	}
	return nums3
}
