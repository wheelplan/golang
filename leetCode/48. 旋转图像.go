//给定一个 n × n 的二维矩阵表示一个图像。
//
//将图像顺时针旋转 90 度。
//
//说明：
//你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
//
//示例 :
//
//给定 matrix =
//[
//[1,2,3],
//[4,5,6],
//[7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//[7,4,1],
//[8,5,2],
//[9,6,3]
//]

package main

import "fmt"

func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	Print(matrix)

	rotate(matrix)

	Print(matrix)

}

func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func Print(matrix [][]int) {

	for i := range matrix {
		fmt.Println(matrix[i])
	}

	fmt.Println()
}

/*func rotate(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)

	for i := range tmp {
		tmp[i] = make([]int, n)
	}

	for i, row := range matrix {
		for j, v := range row {
			tmp[j][n - i - 1] = v
		}
	}

	copy(matrix, tmp)
}*/
