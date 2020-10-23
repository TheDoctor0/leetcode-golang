package solutions

func canConstruct(ransomNote string, magazine string) bool {
    letters := make([]int, 26)

    for _, value := range magazine {
        letters[value - 'a']++
    }

    for _, value := range ransomNote {
        letters[value - 'a']--

        if letters[value - 'a'] < 0 {
            return false
        }
    }

    return true
}
