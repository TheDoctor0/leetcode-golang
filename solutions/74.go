package solutions

func searchMatrix(matrix [][]int, target int) bool {
    row := len(matrix) - 1
    column := 0

    for row >= 0 && column < len(matrix[0]) {
        if matrix[row][column] == target {
            return true
        }

        if matrix[row][column] < target {
            column++

            continue
        }

        row--
    }

    return false
}
