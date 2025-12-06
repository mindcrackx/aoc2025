package utils

func RotateMatrix2DCounterClockwise[T any](matrix [][]T) [][]T {
	rowLength := len(matrix)
	colLength := len(matrix[0])

	resultMatrix := make([][]T, 0, colLength)

	for i := 0; i < colLength; i++ {
		tmp := make([]T, 0, rowLength)
		for j := 0; j < rowLength; j++ {
			tmp = append(tmp, matrix[j][colLength-1-i])
		}
		resultMatrix = append(resultMatrix, tmp)
	}

	return resultMatrix
}
