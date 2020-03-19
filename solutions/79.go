package solutions

func exist(board [][]byte, word string) bool {
    if len(word) == 0 {
        return false
    }

    search := []byte(word)

    for x := range board {
        for y := range board[x] {
            if searchWord(x, y, board, search) {
                return true
            }
        }
    }

    return false
}

func searchWord(x, y int, board [][]byte, word []byte) bool {
    if (x < 0 || x >= len(board)) || (y < 0 || y >= len(board[x])) {
        return false
    }

    if board[x][y] != word[0] {
        return false
    }

    if len(word) == 1 {
        return true
    }

    temp := board[x][y]
    board[x][y] = '0'

    found := searchWord(x + 1, y, board, word[1:]) ||
        searchWord(x - 1, y, board, word[1:]) ||
        searchWord(x, y + 1, board, word[1:]) ||
        searchWord(x, y - 1, board, word[1:])

    board[x][y] = temp

    return found
}
