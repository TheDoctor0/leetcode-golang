package solutions

func isSubsequence(s string, t string) bool {
    currentIndex, maxIndex := 0, len(s)

    if maxIndex == 0 {
        return true
    }

    for i := 0; i < len(t); i++ {
        if t[i] == s[currentIndex] {
            currentIndex++
        }

        if currentIndex == maxIndex {
            return true
        }
    }

    return false
}
