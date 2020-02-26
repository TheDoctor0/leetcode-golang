package solutions

func lengthOfLongestSubstring(s string) int {
    var charPosition [256]int
    var maxLength, substringLength, repeatedPosition int

    for i := 0; i < len(s); i++ {
        if position := charPosition[s[i]]; position > 0 {
            maxLength = max(substringLength, maxLength)
            repeatedPosition = max(position, repeatedPosition)

            substringLength = i + 1 - repeatedPosition
        } else {
            substringLength += 1
        }

        charPosition[s[i]] = i + 1
    }

    return max(maxLength, substringLength)
}
