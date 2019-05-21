package main

func main() {
	rotateArray([]int{1, 2, 3, 4, 5}, 2)
}

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
func rotateArray(nums []int, k int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	length := len(nums)
	offset := k % length
	s1 := make([]int, length-offset)
	s2 := make([]int, offset)
	copy(s1, nums[:length-offset])
	copy(s2, nums[length-offset:])

	copy(nums, append(s2, s1...))
}
