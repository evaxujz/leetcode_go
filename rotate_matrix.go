package main

func main() {

}

// 旋转图像

// 给定一个 n × n 的二维矩阵表示一个图像。
// 将图像顺时针旋转 90 度。
func rotateMatrix(matrix [][]int) {
	size := len(matrix) - 1
	for min, max := 0, size; min < max; min, max = min+1, max-1 {
		round := max - min
		for i := 0; i < round; i++ {
			tmp := matrix[min][min+i]
			matrix[min][min+i] = matrix[max-i][min]
			matrix[max-i][min] = matrix[max][max-i]
			matrix[max][max-i] = matrix[min+i][max]
			matrix[min+i][max] = tmp
		}
	}
}
