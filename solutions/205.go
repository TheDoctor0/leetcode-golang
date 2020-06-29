package solutions

func isIsomorphic(s string, t string) bool {
    first, second := make([]int, 128), make([]int, 128)

    for i := 0; i < len(s); i++ {
        if first[s[i]] != second[t[i]] {
            return false
        }

        first[s[i]], second[t[i]] = i + 1, i + 1
    }

    return true
}