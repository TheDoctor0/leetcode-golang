package solutions

func spiralOrder(matrix [][]int) []int {
    var result []int

    if len(matrix) == 0 {
        return result
    }

    rowBegin := 0
    rowEnd := len(matrix) - 1
    columnBegin := 0
    columnEnd := len(matrix[0]) - 1

    for rowBegin <= rowEnd && columnBegin <= columnEnd {
        for i := columnBegin; i <= columnEnd; i++ {
            result = append(result, matrix[rowBegin][i])
        }

        rowBegin++

        for i := rowBegin; i <= rowEnd; i++ {
            result = append(result, matrix[i][columnEnd])
        }

        columnEnd--

        if rowBegin <= rowEnd {
            for i := columnEnd; i >= columnBegin; i-- {
                result = append(result, matrix[rowEnd][i])
            }
        }

        rowEnd--

        if columnBegin <= columnEnd {
            for i := rowEnd; i >= rowBegin; i-- {
                result = append(result, matrix[i][columnBegin])
            }
        }

        columnBegin++
    }

    return result
}
