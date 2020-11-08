package solutions

import (
    "strconv"
)

func findNthDigit(n int) int {
    table := make([]int, 10)
    table[0] = 9
    current := 9

    for i := 1; i < 10; i ++{
        current = current * 10
        table[i] = table[i - 1] + current * (i + 1)
    }

    length := 0
    current = 1

    for ; length < 10; length++{
        if n <= table[length] {
            break
        }

        current = current * 10
    }

    if length == 0 {
        return n
    }

    last := table[length - 1]
    number := strconv.Itoa((n - last- 1) / (length + 1) + current)

    return int(number[(n - last- 1) % (length + 1)] - '0')
}
