package solutions

func setZeroes(matrix [][]int)  {
    rows := len(matrix)
    columns := len(matrix[0])
    firstColumn := false

    for row := 0; row < rows; row++ {
        if matrix[row][0] == 0 {
            firstColumn = true
        }

        for column := 1; column < columns; column++ {
            if matrix[row][column] == 0 {
                matrix[row][0] = 0
                matrix[0][column] = 0
            }
        }
    }

    for row := 1; row < rows; row++ {
        for column := 1; column < columns; column++ {
            if matrix[0][column] == 0 || matrix[row][0] == 0 {
                matrix[row][column] = 0
            }
        }
    }

    if matrix[0][0] == 0 {
        for column := 0; column < columns; column++ {
            matrix[0][column] = 0
        }
    }

    if firstColumn {
        for row := 0; row < rows; row++ {
            matrix[row][0] = 0
        }
    }
}
