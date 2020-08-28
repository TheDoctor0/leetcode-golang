package solutions

func removeInvalidParentheses(s string) []string {
    var result []string
    left, right := 0, 0

    for _, value := range s {
        if value == '(' {
            left++
        } else if value == ')' {
            if left == 0 {
                right++
            } else {
                left--
            }
        }
    }

    removeParentheses(s, 0, left, right, &result)

    return result
}

func removeParentheses(s string, start, left, right int, result *[]string) {
    if left == 0 && right == 0 {
        if isValidParentheses(s) {
            *result = append(*result, s)
        }

        return
    }

    for i := start; i < len(s); i++ {
        if i != start && s[i] == s[i - 1] {
            continue
        }
        if s[i] == ')' && right > 0 {
            current := s
            current = current[:i] + current[i + 1:]

            removeParentheses(current, i, left, right - 1, result)
        } else if s[i] == '(' && left > 0 {
            current := s
            current = current[:i] + current[i + 1:]

            removeParentheses(current, i, left - 1, right, result)
        }
    }
}

func isValidParentheses(s string) bool {
    count := 0

    for _, value := range s {
        if value == '(' {
            count++
        }

        if value == ')' {
            count--
        }

        if count < 0 {
            return false
        }
    }

    return count == 0
}
