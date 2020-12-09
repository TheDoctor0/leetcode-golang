package solutions

import (
    "fmt"
)

func compress(chars []byte) int {
    scan, write := 0, 0
    length := len(chars)

    for scan < length {
        count := 0
        chars[write] = chars[scan]

        for scan < length && chars[write] == chars[scan] {
            count++
            scan++
        }

        if count > 1 {
            current := fmt.Sprintf("%d", count)

            for _, character := range []byte(current) {
                write++
                chars[write] = character
            }
        }

        write++
    }

    return write
}
