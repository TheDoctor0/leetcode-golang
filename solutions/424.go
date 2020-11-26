package solutions

func characterReplacement(s string, k int) int {
    counts := make([]int, 26)
    maxCount, result := 0, 0

    for i, j := 0, 0; j < len(s); j++ {
        counts[s[j] - 'A']++
        maxCount = max(maxCount, counts[s[j] - 'A'])

        for i <= j && j - i + 1 - maxCount > k {
            counts[s[i] - 'A']--
            i++
        }

        result = max(result, j - i + 1)
    }

    return result
}
