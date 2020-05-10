package solutions

func partition(s string) [][]string {
    var result [][]string

    partitionPalindrome(s, []string{}, &result)

    return result
}

func partitionPalindrome(s string, partition []string, result *[][]string) {
    if len(s) == 0 {
        *result = append(*result, partition)

        return
    }

    for i := 0; i <= len(s); i++ {
        if isValidPalindrome(s[: i]) {
            partitionPalindrome(s[i:], append(append([]string{}, partition...), s[: i]), result)
        }
    }
}

func isValidPalindrome(s string) bool {
    if len(s) == 0 {
        return false
    }

    i, j := 0, len(s) - 1

    for i < j {
        if s[i] != s[j] {
            return false
        }

        i++
        j--
    }

    return true
}
