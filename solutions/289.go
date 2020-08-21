package solutions

func gameOfLife(board [][]int)  {
    if len(board) == 0 || len(board[0]) == 0 {
        return
    }

    x, y := len(board), len(board[0])

    for i := 0; i < x; i++ {
        for j := 0; j < y; j++ {
            lives := 0
            lives += isAlive(board, i - 1, j - 1)
            lives += isAlive(board, i - 1, j)
            lives += isAlive(board, i - 1, j + 1)
            lives += isAlive(board, i, j - 1)
            lives += isAlive(board, i, j + 1)
            lives += isAlive(board, i + 1, j - 1)
            lives += isAlive(board, i + 1, j)
            lives += isAlive(board, i + 1, j + 1)

            if board[i][j] == 0 {
                if lives == 3 {
                    board[i][j] = 2
                }
            } else {
                if lives < 2 || lives > 3 {
                    board[i][j] = 3
                }
            }
        }
    }

    for i := 0; i < x; i++ {
        for j := 0; j < y; j++ {
            if board[i][j] == 2 {
                board[i][j] = 1
            } else if board[i][j] == 3 {
                board[i][j] = 0
            }
        }
    }
}

func isAlive(board [][]int, i int, j int) int {
    if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 0 || board[i][j] == 2 {
        return 0
    }

    return 1
}
