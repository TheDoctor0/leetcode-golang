package solutions

func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0' {
        return 0
    }

    result := make([]int, len(s) + 1)
    result[0] = 1

    for i := 1; i <= len(s); i++ {
        if s[i - 1] != '0' {
            result[i] = result[i - 1]
        }

        if i >= 2 && (s[i - 2] == '1' || (s[i - 2] == '2' && s[i - 1] <= '6')) {
            result[i] += result[i - 2]
        }
    }

    return result[len(s)]
}