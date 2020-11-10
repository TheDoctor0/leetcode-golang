package solutions

func removeKdigits(num string, k int) string {
    count := k

    var stack []byte
    var result string

    for i := 0; i < len(num); i++ {
        for k >= 0 {
            if k == 0 {
                result = string(stack) + num[i:]
            }

            if len(stack) == 0  {
                stack = append(stack, num[i])
                break
            } else if stack[len(stack) - 1] <= num[i] {
                stack = append(stack, num[i])
                break
            } else {
                stack = stack[:len(stack) - 1]
                k--
            }
        }
    }

    if k > 0 {
        result = string(stack)[:len(num) - count]
    }

    for i := 0; i < len(result); i++ {
        if result[i] != '0' {
            return result[i:]
        }
    }

    return "0"
}
