package solutions

func getRow(rowIndex int) []int {
    result := make([][]int, 2)

    for i := 0; i < 2; i++ {
        result[i] = make([]int, rowIndex+1)
    }

    result[1][0] = 1

    for i := 0; i < rowIndex; i++ {
        result[i % 2][0], result[i % 2][i + 1] = 1, 1

        for j := 1; j <= i; j++ {
            result[i % 2][j] = result[1 - i % 2][j] + result[1 - i % 2][j - 1]
        }
    }

    return result[1 - rowIndex % 2]
}
