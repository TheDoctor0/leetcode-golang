package solutions

import (
    "strconv"
)

type operation func(firstOperand, secondOperand int) int

func addOperator(firstOperand, secondOperand int) int {
    return firstOperand + secondOperand
}

func subtractOperator(firstOperand, secondOperand int) int {
    return firstOperand - secondOperand
}

func multiplyOperator(firstOperand, secondOperand int) int {
    return firstOperand * secondOperand
}

func divideOperator(firstOperand, secondOperand int) int {
    return firstOperand / secondOperand
}

func evalRPN(tokens []string) int {
    var cursor int
    var operand []int
    var operators = map[string]operation {
        "+": addOperator,
        "-": subtractOperator,
        "*": multiplyOperator,
        "/": divideOperator,
    }

    for _, token := range tokens {
        if function, ok := operators[token]; ok {
            firstOperand, secondOperand := operand[cursor - 2], operand[cursor - 1]
            result := function(firstOperand, secondOperand)

            operand[cursor - 2] = result
            operand = operand[: cursor - 1]
            cursor--

            continue
        }

        value, _ := strconv.Atoi(token)
        operand = append(operand, value)
        cursor++
    }

    return operand[0]
}
