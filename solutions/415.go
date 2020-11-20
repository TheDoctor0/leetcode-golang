package solutions

func addStrings(num1 string, num2 string) string {
    firstString, secondString := []byte(num1), []byte(num2)

    if len(firstString) < len(secondString) {
        firstString, secondString = secondString, firstString
    }

    sum := byte(0)

    for i, j := len(firstString) - 1, len(secondString) - 1; i >= 0 || j >= 0; i, sum = i - 1, sum / 10 {
        if j >= 0 {
            sum += secondString[j] - '0'
            j--
        }

        sum += firstString[i] - '0'
        firstString[i] = sum % 10 + '0'
    }

    if (sum) != byte(0) {
        firstString = append([]byte{'1'}, firstString...)
    }

    return string(firstString)
}
