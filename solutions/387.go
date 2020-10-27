package solutions

func firstUniqChar(s string) int {
    letters := make([]int, 26)

    for _, character := range s {
        letters[character - 'a']++
    }

    for i, character := range s {
        if letters[character - 'a'] == 1 {
            return i
        }
    }

    return -1
}
