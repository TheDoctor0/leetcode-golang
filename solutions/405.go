package solutions

func toHex(num int) string {
    if num == 0 {
        return "0"
    }

    hex := "0123456789abcdef"
    sum, result := 0, ""

    for i := 31; i >= 0; i-- {
        sum = sum << 1 + (num >> i) & 1

        if i % 4 == 0 && !(len(result) == 0 && sum == 0) {
            result += string(hex[sum])
            sum = 0
        }
    }

    return result
}
