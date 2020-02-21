package solutions

func totalNQueens(n int) int {
    count := 0

    backtrack(&count, n, 0, 0, 0, 0)

    return count
}

func backtrack(count *int, n int, row int, column int, left int, right int) {
    if row == n {
        *count++

        return
    }

    for i := 0; i < n; i++ {
        bit := 1 << uint(n - i - 1)

        if column & bit == 0 && left & bit == 0 && right & bit == 0 {
            backtrack(count, n, row + 1, column | bit, (left | bit) << 1, (right | bit) >> 1)
        }
    }
}
