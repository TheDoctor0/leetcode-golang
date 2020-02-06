package solutions

type Sudoku struct {
    board [][]byte
    rows [90]bool
    columns [90]bool
    squares [230]bool
    solved bool
}

func solveSudoku(board [][]byte) {
    s := createBoard(board)

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if s.board[i][j] == '.' {
                s.backtrack(i, j)

                if s.solved {
                    return
                }
            }
        }
    }
}

func createBoard(board [][]byte) *Sudoku {
    s := &Sudoku{
        board: board,
    }

    for i, row := range board {
        for j, value := range row {
            if value != '.' {
                s.insert(i, j, value)
            }
        }
    }

    return s
}

func (s *Sudoku) canInsert(x, y int, input byte) bool {
    if s.board[x][y] != '.' {
        return false
    }

    value := int(input - '0')

    if s.rows[s.axisIndex(x, value)] || s.columns[s.axisIndex(y, value)] || s.squares[s.squareIndex(x, y, value)] {
        return false
    }

    return true
}

func (s *Sudoku) insert(x, y int, input byte) {
    s.board[x][y] = input

    value := int(input - '0')

    s.rows[s.axisIndex(x, value)] = true
    s.columns[s.axisIndex(y, value)] = true
    s.squares[s.squareIndex(x, y, value)] = true
}

func (s *Sudoku) remove(x, y int, input byte) {
    s.board[x][y] = '.'

    value := int(input - '0')

    s.rows[s.axisIndex(x, value)] = false
    s.columns[s.axisIndex(y, value)] = false
    s.squares[s.squareIndex(x, y, value)] = false
}

func (s *Sudoku) axisIndex(axis, value int) int {
    return axis * 10 + value
}

func (s *Sudoku) squareIndex(x, y, value int) int {
    return (x / 3) * 100 + (y / 3) * 10 + value
}

func (s *Sudoku) backtrack(x, y int) {
    if y == 9 {
        x = x + 1
        y = 0
    }
    
    if x == 9 {
        s.solved = true

        return
    }

    if s.board[x][y] == '.' {
        for _, option := range []byte {'1', '2', '3', '4', '5', '6', '7', '8', '9'} {
            if !s.canInsert(x, y, option) {
                continue
            }

            s.insert(x, y, option)
            s.backtrack(x, y + 1)

            if s.solved {
                return
            }

            s.remove(x, y, option)
        }
    } else {
        s.backtrack(x, y + 1)
    }
}
