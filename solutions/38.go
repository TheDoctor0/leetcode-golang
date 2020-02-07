package solutions

import (
    "strconv"
)

func countAndSay(n int) string {
    result := "1"

    if n < 1 {
        return result
    }

    for i := 1; i < n; i++ {
        sequence := make([]byte, 0)
        count := 1

        for i := 1; i < len(result); i++ {
            if result[i] == result[i - 1] {
                count++
            } else {
                sequence = append(sequence, []byte(strconv.Itoa(count))...)
                sequence = append(sequence, result[i - 1])
                count = 1
            }
        }

        sequence = append(sequence, []byte(strconv.Itoa(count))...)
        sequence = append(sequence, result[len(result) - 1])

        result = string(sequence)
    }

    return result
}
