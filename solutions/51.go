package solutions

func solveNQueens(n int) [][]string {
    var result = make([][]string, 0)

    backtrackBoard(&result, makeEmptyBoard(n), n, 0, 0, 0, 0)

    return result
}

func makeEmptyBoard(n int) [][]byte {
    board := make([][]byte, n)

    for i := 0; i < n; i++ {
        board[i] = make([]byte, n)

        for j := 0; j < n; j++ {
            board[i][j] = '.'
        }
    }

    return board
}

func backtrackBoard(result *[][]string, board [][]byte, n int, row int, column int, left int, right int) {
    if row == n {
        temp := make([]string, n)

        for i, value := range board {
            temp[i] = string(value)
        }

        *result = append(*result, temp)

        return
    }

    for i := 0; i < n; i++ {
        bit := 1 << uint(n - i - 1)

        if column & bit == 0 && left & bit == 0 && right & bit == 0 {
            board[row][i] = 'Q'

            backtrackBoard(result, board, n, row + 1, column | bit, (left | bit) << 1, (right | bit) >> 1)

            board[row][i] = '.'
        }
    }
}
