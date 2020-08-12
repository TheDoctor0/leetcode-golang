package solutions

import (
    "strings"
)

func numberToWords(num int) string {
    const (
        billion  = 1000000000
        million  = 1000000
        thousand = 1000
        hundred  = 100
        ten      = 10
    )

    numbers := map[int]string{
        0:  "Zero",
        1:  "One",
        2:  "Two",
        3:  "Three",
        4:  "Four",
        5:  "Five",
        6:  "Six",
        7:  "Seven",
        8:  "Eight",
        9:  "Nine",
        10: "Ten",
        11: "Eleven",
        12: "Twelve",
        13: "Thirteen",
        14: "Fourteen",
        15: "Fifteen",
        16: "Sixteen",
        17: "Seventeen",
        18: "Eighteen",
        19: "Nineteen",
        20: "Twenty",
        30: "Thirty",
        40: "Forty",
        50: "Fifty",
        60: "Sixty",
        70: "Seventy",
        80: "Eighty",
        90: "Ninety",
    }

    if value, ok := numbers[num]; ok {
        return value
    }

    var result []string

    base := num / billion
    num = num % billion

    if base > 0 {
        result = append(result, numberToWords(base))
        result = append(result, "Billion")
    }

    base = num / million
    num = num % million

    if base > 0 {
        result = append(result, numberToWords(base))
        result = append(result, "Million")
    }

    base = num / thousand
    num = num % thousand

    if base > 0 {
        result = append(result, numberToWords(base))
        result = append(result, "Thousand")
    }

    base = num / hundred
    num = num % hundred

    if base > 0 {
        result = append(result, numberToWords(base))
        result = append(result, "Hundred")
    }

    if num >= 20 {
        base = num / ten
        num = num % ten

        result = append(result, numberToWords(base * ten))
    }

    if num > 0 {
        result = append(result, numberToWords(num))
    }

    return strings.Join(result, " ")
}
