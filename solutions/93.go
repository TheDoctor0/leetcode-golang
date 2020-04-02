package solutions

import (
    "strconv"
    "strings"
)

func restoreIpAddresses(s string) []string {
    result := make([]string, 0)

    backtrackIpAddresses(s, &result, nil)

    return result
}

func backtrackIpAddresses(s string, result *[]string, current []string) {
    if len(s) == 0 && len(current) == 4 {
        *result = append(*result, strings.Join(current, "."))

        return
    }

    if len(current) > 4 {
        return
    }

    for i := 1; i <= 3 && i <= len(s); i++ {
        part := s[0: i]

        if strings.HasPrefix(part, "0") && len(part) != 1 {
            continue
        }

        if number, _ := strconv.Atoi(part); number > 255 {
            continue
        }

        backtrackIpAddresses(s[i:], result, append(current, part))
    }
}
