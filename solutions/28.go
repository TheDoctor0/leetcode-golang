package solutions

func strStr(haystack string, needle string) int {
    if len(needle) == 0 || haystack == needle {
        return 0
    }

    if len(haystack) == 0 {
        return -1
    }

    for i := 0; i < (len(haystack) - len(needle) + 1); i++ {
        word := haystack[i:len(needle) + i]

        if (word) == needle{
            return i
        }
    }

    return -1
}
