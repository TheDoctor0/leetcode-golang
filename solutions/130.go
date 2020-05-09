package solutions

func solve(board [][]byte)  {
    if len(board) == 0 {
        return
    }

    for i := 0; i < len(board); i++ {
        walkThroughBoard(i, 0, board)
        walkThroughBoard(i, len(board[0]) - 1, board)
    }

    for j := 0; j < len(board[0]); j++ {
        walkThroughBoard(0, j, board)
        walkThroughBoard(len(board) - 1, j, board)
    }

    for i := range board {
        for j := range board[i] {
            if board[i][j] == 'S' {
                board[i][j] = 'O'
            } else {
                board[i][j] = 'X'
            }
        }
    }
}

func walkThroughBoard(i, j int, board [][]byte) {
    if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
        return
    }

    if board[i][j] == 'O' {
        board[i][j] = 'S'

        walkThroughBoard(i - 1, j, board)
        walkThroughBoard(i + 1, j, board)
        walkThroughBoard(i, j - 1, board)
        walkThroughBoard(i,  j + 1, board)
    }
}
