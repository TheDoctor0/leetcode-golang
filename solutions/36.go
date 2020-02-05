package solutions

func isValidSudoku(board [][]byte) bool {
    return blocksValid(board) && columnsValid(board) && rowsValid(board)
}

func columnsValid(board [][]byte) bool {
    for column := 0; column < 9; column++ {
        set := make([]int, 10)

        for i := 0; i < 9; i++ {
            if !cellValid(set, board[i][column]) {
                return false
            }
        }
    }

    return true
}

func rowsValid(board [][]byte) bool {
    for row := 0; row < 9; row++ {
        set := make([]int, 10)

        for i := 0; i < 9; i++ {
            if !cellValid(set, board[row][i]) {
                return false
            }
        }
    }

    return true
}

func blocksValid(board [][]byte) bool {
    for row := 0; row < 9; row = row + 3 {
        for column := 0; column < 9; column = column + 3 {
            set := make([]int, 10)

            for i := row; i < (row + 3); i++ {
                for j := column; j < (column + 3); j++ {
                    if !cellValid(set, board[i][j]) {
                        return false
                    }
                }
            }
        }
    }

    return true
}

func cellValid(set []int, cellVal byte) bool {
    if cellVal == '.' {
        return true
    }

    if set[cellVal - 48] == 1 {
        return false
    }

    set[cellVal - 48] = 1

    return true
}
