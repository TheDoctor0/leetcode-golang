package solutions

import (
    "strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 || denominator == 0 {
        return "0"
    }

    sign := 1

    if numerator < 0 {
        sign *= -1
        numerator *= -1
    }

    if denominator < 0 {
        sign *= -1
        denominator *= -1
    }

    signPrefix := ""

    if sign < 0 {
        signPrefix = "-"
    }

    integerPart, decimalPart := numerator / denominator, numerator % denominator
    integerPartInString := strconv.Itoa(integerPart)

    if decimalPart == 0 {
        return signPrefix + integerPartInString
    }

    order, currentDecimal := 1, decimalPart
    decimalPartInBytes, decimalOrder := make([]byte, 0), make(map[int]int)

    for currentDecimal != 0 {
        if decimalOrder[currentDecimal] != 0 {
            repeatHeadPosition := decimalOrder[currentDecimal] - 1
            repeatPart := make([]byte, order - decimalOrder[currentDecimal])

            copy(repeatPart, decimalPartInBytes[repeatHeadPosition:])

            decimalPartInBytes = decimalPartInBytes[: repeatHeadPosition]
            decimalPartInBytes = append(decimalPartInBytes, '(')
            decimalPartInBytes = append(decimalPartInBytes, repeatPart...)
            decimalPartInBytes = append(decimalPartInBytes, ')')

            break
        }

        decimalOrder[currentDecimal] = order
        decimalPartInBytes = append(decimalPartInBytes, byte(currentDecimal * 10 / denominator) + '0')
        currentDecimal = currentDecimal * 10 % denominator

        order++
    }

    return signPrefix + integerPartInString + "." + string(decimalPartInBytes)
}
