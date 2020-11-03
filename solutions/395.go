package solutions

func longestSubstring(s string, k int) int {
    dictionary := make([]int, 26)

    for _, character := range s {
        dictionary[int(character - 'a')]++
    }

    result, temp := "", ""

    for i := 0; i < len(s); i++ {
        if dictionary[int(s[i] - 'a')] < k {
            result = maxLength(result, temp)
            temp = ""
        } else {
            temp += string(s[i])
        }
    }

    result = maxLength(result, temp)

    if len(result) == len(s) {
        return len(result)
    }

    return longestSubstring(result, k)
}

func maxLength(a string, b string) string {
    if len(a) > len(b) {
        return a
    }

    return b
}
