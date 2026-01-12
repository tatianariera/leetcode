package leetcode

func rotateMatrix(matrix [][]int) {

	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	for i := 0; i < n; i++ {
		left, right := 0, n-1

		for left < right {
			temp := matrix[i][left]
			matrix[i][left] = matrix[i][right]
			matrix[i][right] = temp
			left++
			right--
		}
	}
}
