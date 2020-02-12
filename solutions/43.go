package solutions

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    result := make([]byte, len(num1) + len(num2))

    for i := len(num2) - 1; i >= 0; i-- {
        for j := len(num1) - 1; j >= 0; j-- {
            current := (num2[i] - '0') * (num1[j] - '0') + result[i + j + 1]

            result[i + j + 1] = current % 10
            result[i + j] += current / 10
        }
    }

    var zero int

    for i := 0; i < len(result) - 1; i++ {
        if result[i] != 0 {
            break
        }

        zero++
    }

    for i := zero; i < len(result); i++ {
        result[i] += '0'
    }

    return string(result[zero:])
}
