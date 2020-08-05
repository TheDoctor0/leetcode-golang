package solutions

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    alphabet := make([]int, 26)

    for _, char := range s {
        alphabet[int(char - 'a')]++
    }

    for _, char := range t {
        if alphabet[int(char - 'a')] <= 0 {
            return false
        }

        alphabet[int(char - 'a')]--
    }

    return true
}
