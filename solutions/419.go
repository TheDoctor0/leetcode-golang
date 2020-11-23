package solutions

func countBattleships(board [][]byte) int {
    result := 0

    for i, array := range board {
        for j, value := range array {
            if value == 'X' {
                switch {
                case i == 0 && j == 0:
                    result++
                case i != 0 && j == 0:
                    if board[i - 1][j] != 'X' {
                        result++
                    }
                case i == 0 && j != 0:
                    if board[i][j - 1] != 'X' {
                        result++
                    }
                case i != 0 && j != 0:
                    if board[i][j - 1] != 'X' && board[i - 1][j] != 'X' {
                        result++
                    }
                }
            }
        }
    }

    return result
}
