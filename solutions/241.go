package solutions

import (
    "strconv"
)

func diffWaysToCompute(input string) []int {
    var result []int

    for i := 0; i < len(input); i++ {
        if input[i] == '+' || input[i] == '-' || input[i] == '*' {
            for _, a := range diffWaysToCompute(input[:i]) {
                for _, b := range diffWaysToCompute(input[i + 1:]) {
                    switch input[i] {
                    case '+':
                        result = append(result, a + b)
                    case '-':
                        result = append(result, a - b)
                    case '*':
                        result = append(result, a * b)
                    }
                }
            }
        }
    }

    if value, _ := strconv.Atoi(input); len(result) == 0 {
        result = append(result, value)
    }

    return result
}
