package solutions

func longestCommonPrefix(strs []string) string {
    if strs == nil || len(strs) == 0 {
        return ""
    }

    first := strs[0]
    fromSecond := strs[1:]

    for i := 0; i < len(first); i++ {
        letter := first[i]

        for _, s := range fromSecond {
            if len(s) <= i || letter != s[i] {
                return first[:i]
            }
        }
    }

    return first;
}
