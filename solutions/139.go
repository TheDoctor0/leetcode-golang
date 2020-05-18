package solutions

func wordBreak(s string, wordDict []string) bool {
    result := make([]bool, len(s) + 1)
    result[0] = true

    for i := 0 ; i < len(s); i++ {
        if result[i] == false {
            continue
        }

        for _, word := range wordDict {
            j := i + len(word)

            if j <= len(s) && s[i: j] == word {
                result[j] = true
            }
        }
    }

    return result[len(s)]
}
