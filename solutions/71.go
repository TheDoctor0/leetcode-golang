package solutions

import (
    "strings"
)

func simplifyPath(path string) string {
    var result []string
    parts := strings.Split(path, "/")

    for _, part := range parts {
        if part == "" || part == "." {
            continue
        }

        if part == ".." {
            if len(result) > 0 {
                result = result[:len(result) - 1]
            }

            continue
        }

        result = append(result, part)
    }

    return "/" + strings.Join(result, "/")
}
