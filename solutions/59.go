package solutions

func generateMatrix(n int) [][]int {
    matrix := make([][]int, n)

    for i := 0; i < n; i++ {
        matrix[i] = make([]int, n)
    }

    top, right, bottom, left := 0, len(matrix[0]) - 1, len(matrix) - 1, 0

    for number := 1; top <= bottom && left <= right; {
        for i := left; i <= right; i++ {
            matrix[top][i] = number
            number++
        }

        top++

        for i := top; i <= bottom; i++ {
            matrix[i][right] = number
            number++
        }

        right--

        for i := right; i >= left; i-- {
            matrix[bottom][i] = number
            number++
        }

        bottom--

        for i := bottom; i >= top; i-- {
            matrix[i][left] = number
            number++
        }

        left++
    }

    return matrix
}
