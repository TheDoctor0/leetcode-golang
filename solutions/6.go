package solutions

import (
    "bytes"
)

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    var result bytes.Buffer

    for i := 0; i < numRows; i++ {
        for j := i; j < len(s); {
            result.WriteByte(s[j])

            j += numRows * 2 - 2

            if i == 0 || i == numRows - 1 {
                continue
            }

            j -= i + i

            if j >= len(s) {
                break
            }

            result.WriteByte(s[j])

            j += i + i
        }
    }

    return result.String()
}
