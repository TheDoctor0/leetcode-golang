package solutions

func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }

    row, column := len(matrix), len(matrix[0])
    previous, maxSize := 0, 0
    current := make([]int, column)

    for i := 0; i < row; i++ {
        for j := 0; j < column; j++ {
            temp := current[j]

            if i == 0 || j == 0 || matrix[i][j] == '0' {
                current[j] = int(matrix[i][j] - '0')
            } else {
                current[j] = min(previous, min(current[j], current[j - 1])) + 1
            }

            if maxSize < current[j] {
                maxSize = current[j]
            }

            previous = temp
        }
    }

    return maxSize * maxSize
}
