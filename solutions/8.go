package solutions

import (
    "math"
)

func myAtoi(str string) int {
    sign := 1
    integer := 0
    conversion := false

    for i := 0; i < len(str); i++ {
        if str[i] == ' ' && !conversion {
            continue
        }

        if !conversion {
            if str[i] == '+' {
                conversion = true
            } else if str[i] == '-' {
                conversion = true
                sign = -1
            } else if digit := toDigit(str[i]); digit > -1 {
                integer = integer * 10 + digit
                conversion = true
            } else {
                return 0
            }
        } else {
            if digit := toDigit(str[i]); digit > -1 {
                integer = integer * 10 + digit

                if integer * sign <= math.MinInt32 {
                    return math.MinInt32
                } else if integer * sign >= math.MaxInt32 {
                    return math.MaxInt32
                }
            } else {
                break
            }
        }
    }

    return integer * sign
}

func toDigit(char byte) int {
    if '0' <= char && char <= '9' {
        return int(char - '0')
    }

    return -1
}