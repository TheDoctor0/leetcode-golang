package solutions

func wordBreak(s string, wordDict []string) []string {
    return findWordBreak(s, wordDict, map[string][]string{})
}

func findWordBreak(s string, wordsList []string, cache map[string][]string) []string {
    if result, ok := cache[s]; ok {
        return result
    }

    if len(s) == 0 {
        return []string{""}
    }

    var result []string

    for _, word := range wordsList {
        if len(word) <= len(s) && word == s[: len(word)] {
            for _, current := range findWordBreak(s[len(word):], wordsList, cache) {
                if len(current) == 0 {
                    result = append(result, word)
                } else {
                    result = append(result, word + " " +current)
                }
            }
        }
    }

    cache[s] = result

    return result
}
