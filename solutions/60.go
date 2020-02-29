package solutions

import (
    "strconv"
)

func getPermutation(n int, k int) string {
    list := make([]string, n)
    factorials := make([]int, n + 1)
    factorials[0] = 1

    for i := 1; i <= n; i++ {
        list[i - 1] = strconv.Itoa(i)

        factorials[i] = i * factorials[i - 1]
    }

    result := ""
    k = k - 1

    for i := 1; i <= n; i++ {
        position := k / factorials[n - i]

        result = result + list[position]
        list = append(list[:position], list[position + 1:]...)

        k = k % factorials[n - i]
    }

    return result
}
