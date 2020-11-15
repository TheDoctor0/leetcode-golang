package solutions

func longestPalindrome(s string) int {
    cache := make(map[byte]struct{})

    for i := range s {
        if _, ok := cache[s[i]]; ok {
            delete(cache, s[i])
        } else {
            cache[s[i]] = struct{}{}
        }
    }

    if len(cache) == 0 {
        return len(s)
    }

    return len(s) - len(cache) + 1
}
