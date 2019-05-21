package main

func main() {

}

// 移动零

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
func moveZeroes(nums []int) {
	anchor := 0

	for _, value := range nums {
		if value != 0 {
			nums[anchor] = value
			anchor += 1
		}
	}

	for i := anchor; i < len(nums); i++ {
		nums[i] = 0
	}
}
