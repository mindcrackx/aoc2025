package utils

func RotateMatrix2DCounterClockwise[T any](matrix [][]T) [][]T {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]T{}
	}

	rowLength := len(matrix)
	colLength := len(matrix[0])

	resultMatrix := make([][]T, colLength)

	for i := 0; i < colLength; i++ {
		resultMatrix[i] = make([]T, rowLength)
		for j := 0; j < rowLength; j++ {
			resultMatrix[i][j] = matrix[j][colLength-1-i]
		}
	}

	return resultMatrix
}
