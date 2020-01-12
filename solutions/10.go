package solutions

type data struct {
    i int
    j int
}

func isMatch(s string, p string) bool {
    next := make([]data, 0, 10)
    i := len(s) - 1
    j := len(p) - 1

    for i >= 0 || j >= 0 {
        if j >= 0 {
            char := p[j]

            switch char {
            case '.':
                if i >= 0 {
                    j--
                    i--

                    continue
                }
            case '*':
                if i == -1 || (p[j - 1] != '.' && p[j - 1] != s[i]) {
                    j -= 2

                    continue
                }

                next = append(next, data{i - 1, j})
                j -= 2

                continue
            default:
                if i >= 0 && p[j] == s[i] {
                    i--
                    j--

                    continue
                }
            }
        }

        if len(next) != 0 {
            i = next[len(next) - 1].i
            j = next[len(next) - 1].j
            next = next[:len(next) - 1]

            continue
        }

        return false
    }

    return true
}