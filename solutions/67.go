package solutions

import (
    "strconv"
)

func addBinary(a string, b string) string {
    if len(a) < len(b) {
        a, b = b, a
    }

    carry := 0
    result := ""
    zeroes := ""

    for i := 0; i < len(a)-len(b); i++ {
        zeroes += "0"
    }

    b = zeroes + b

    for i := len(a) - 1; i >= 0; i-- {
        sum := carry

        if string(a[i]) == "1" {
            sum++
        }

        if string(b[i]) == "1" {
            sum++
        }

        carry = sum / 2
        sum = sum % 2
        result = strconv.Itoa(sum) + result
    }

    if carry > 0 {
        result = "1" + result
    }

    return result
}
