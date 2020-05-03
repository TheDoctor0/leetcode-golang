package solutions

func numDistinct(s string, t string) int {
    if len(s) < len(t) {
        return 0
    }

    if len(t) == 0 {
        return 1
    }

    result := make([]int, len(t))

    for i := range s {
        previous := 1

        for j := range t {
            current := result[j]

            if t[j] == s[i] {
                result[j] += previous
            }

            previous = current
        }
    }

    return result[len(t) - 1]
}
