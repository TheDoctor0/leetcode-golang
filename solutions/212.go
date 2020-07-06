package solutions

type WordTree struct {
    Nodes [26]*WordTree
    Word string
}

func findWords(board [][]byte, words []string) []string {
    var result []string

    root := buildTrie(words)

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            findWord(board, root, i, j, &result)
        }
    }

    return result
}

func buildTrie(words []string) *WordTree {
    root := &WordTree{}

    for _, word := range words {
        current := root

        for _, c := range word {
            i := c - 'a'

            if current.Nodes[i] == nil {
                current.Nodes[i] = &WordTree{}
            }

            current = current.Nodes[i]
        }

        current.Word = word
    }

    return root
}

func findWord(board [][]byte, root *WordTree, i, j int, result *[]string) {
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
        return
    }

    current := board[i][j]

    if current == '$' || root.Nodes[current - 'a'] == nil {
        return
    }

    root = root.Nodes[current - 'a']

    if len(root.Word) > 0 {
        *result = append(*result, root.Word)

        root.Word = ""
    }

    board[i][j] = '$'

    findWord(board, root, i + 1, j, result)
    findWord(board, root, i - 1, j, result)
    findWord(board, root, i, j + 1, result)
    findWord(board, root, i, j - 1, result)

    board[i][j] = current
}
