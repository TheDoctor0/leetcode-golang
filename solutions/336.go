package solutions

type TrieNode struct {
    nodes [26]*TrieNode
    index int
    match map[int]bool
}

func palindromePairs(words []string) [][]int {
    var result [][]int
    root := &TrieNode{index: -1, match: make(map[int]bool)}

    for i, word := range words {
        addWord(root, word, i)
    }

    for i, word := range words {
        matchSet := searchWord(root, word)

        for matchIndex := range matchSet {
            if matchIndex != i {
                result = append(result, []int{matchIndex, i})
            }
        }
    }

    return result
}

func searchWord(root *TrieNode, word string) map[int]bool {
    result := make(map[int]bool)

    for i := len(word) - 1; i >= 0; i-- {
        if root.index != -1 && isPalindrome(word[: i + 1]) {
            result[root.index] = true
        }

        offset := word[i] - 'a'

        if root.nodes[offset] == nil {
            return result
        }

        root = root.nodes[offset]
    }

    for k := range root.match {
        result[k] = true
    }

    return result
}

func addWord(root *TrieNode, word string, idx int) {
    for i := 0; i < len(word); i++ {
        if isPalindrome(word[i:]) {
            root.match[idx] = true
        }

        offset := word[i] - 'a'

        if root.nodes[offset] == nil {
            root.nodes[offset] = &TrieNode{index: - 1, match: make(map[int]bool)}
        }

        root = root.nodes[offset]
    }

    root.index = idx
    root.match[idx] = true
}

func isPalindrome(s string) bool {
    for left, right := 0, len(s) - 1; left < right; left, right = left + 1, right - 1 {
        if s[left] != s[right] {
            return false
        }
    }

    return true
}
