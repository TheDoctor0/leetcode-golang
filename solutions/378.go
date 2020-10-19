package solutions

func kthSmallest(matrix [][]int, k int) int {
    rows, columns := len(matrix), len(matrix[0])
    low, high := matrix[0][0], matrix[rows - 1][columns - 1] + 1

    for low < high {
        middle := low + (high - low) / 2
        count, i, j := 0, 0, columns - 1

        for ; i < rows; i++ {
            for j >= 0 && middle < matrix[i][j] {
                j--
            }

            count += j + 1
        }

        if count < k {
            low = middle + 1
        } else {
            high = middle
        }
    }

    return low
}
