package solutions

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }

    row, column := len(matrix) - 1, 0

    for {
        if !isInbounds(matrix, row, column) {
            return false
        }

        if value := matrix[row][column]; value == target {
            return true
        } else if value < target {
            column += 1
        } else if value > target {
            row -= 1
        } else {
            return false
        }
    }
}

func isInbounds(matrix [][]int, row int, column int) bool {
    return row >= 0 && column < len(matrix[row])
}
