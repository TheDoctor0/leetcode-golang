package solutions

func calculate(s string) int {
    number, sum, currentSum := 0, 0, 0
    lastSign := byte('+')

    for i := 0; i < len(s); i++ {
        character := s[i]

        if isByteDigit(character) {
            number = 10 * number + int(character - '0')
        }

        if i == len(s) - 1 || (!isByteDigit(character) && character != ' ') {
            switch lastSign {
            case '+':
                sum += currentSum
                currentSum = number
            case '-':
                sum += currentSum
                currentSum = -number
            case '*':
                currentSum *= number
            case '/':
                currentSum /= number
            }

            lastSign = character
            number = 0
        }
    }

    return sum + currentSum
}

func isByteDigit(b byte) bool {
    if b - '0' <= 9 {
        return true
    }

    return false
}
