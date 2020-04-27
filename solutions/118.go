package solutions

func generate(numRows int) [][]int {
    result := make([][]int,numRows)

    for i := 0; i < numRows; i += 1 {
        level := make([]int,i + 1)
        level[0],level[i] = 1, 1

        if i > 1 {
            for j := 1; j <= i / 2; j++ {
                level[j] = result[i - 1][j - 1] + result[i - 1][j]
                level[len(level) - 1 - j] = level[j]
            }
        }

        result[i] = level
    }

    return result
}