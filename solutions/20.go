package solutions

func isValid(s string) bool {
    if s == "" {
        return true
    }

    stack := make([]rune, 0, 1)
    brackets := map[rune]rune{
        '}': '{',
        ')': '(',
        ']': '[',
    }

    for _, char := range s {
        switch char {
        case '(', '[', '{':
            stack = append(stack, char)
        case ')', ']', '}':
            if len(stack) == 0 || stack[len(stack) - 1] != brackets[char] {
                return false
            } else {
                stack = stack[:len(stack) - 1]
            }
        }
    }

    return len(stack) == 0
}
