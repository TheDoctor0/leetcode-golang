package solutions

func reverseVowels(s string) string {
    result := []byte(s)

    for i, j := 0, len(s) - 1; i < len(s) && j > i ; i++ {
        if isVowel(s[i]) {
            for j >= i && !isVowel(s[j]) {
                j--
            }

            result[i], result[j] = result[j], result[i]
            j--
        }
    }

    return string(result)
}

func isVowel(c byte) bool {
    return c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' || c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
