package solutions

func lengthOfLongestSubstring(s string) int {
    var charPosition [256]int
    var maxLength, substringLength, repeatedPosition int

    for i := 0; i < len(s); i++ {
        if position := charPosition[s[i]]; position > 0 {
            maxLength = max(substringLength, maxLength)
            repeatedPosition = max(position, repeatedPosition)

            charPosition[s[i]] = i + 1
            substringLength = i + 1 - repeatedPosition
        } else {
            charPosition[s[i]] = i + 1
            substringLength += 1
        }
    }

    return max(maxLength, substringLength)
}

func max(a int, b int) int {
    if a > b {
        return a
    }

    return b
}
