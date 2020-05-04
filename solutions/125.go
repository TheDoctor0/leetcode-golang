package solutions

func isPalindrome(s string) bool {
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
        for ; i < j && !isChar(s[i]); {
            i++
        }

        for ; i < j && !isChar(s[j]); {
            j--
        }

        if toLower(s[i]) != toLower(s[j]) {
            return false
        }
    }

    return true
}
