package solutions

func rotate(matrix [][]int)  {
    n := len(matrix) - 1

    for row := 0; row < (n + 1) / 2; row++ {
        for column := row; column < n - row; column++ {
            temp := matrix[row][column]
            matrix[row][column] = matrix[n - column][row]
            matrix[n - column][row] = matrix[n - row][n - column]
            matrix[n - row][n - column] = matrix[column][n - row]
            matrix[column][n - row] = temp
        }
    }
}
