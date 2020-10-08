package solutions

import (
    "math"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
    matrixCopy := matrix
    rows, columns := len(matrixCopy), len(matrixCopy[0])
    result := math.MinInt64

    for i := 0; i < rows; i++ {
        list := make([]int, columns)

        for row := i; row < rows; row++ {
            for column := 0; column < columns; column++ {
                list[column] += matrixCopy[row][column]
            }

            current := maxSumArrayWithK(list, k)

            if current == k {
                return k
            }

            if current < k && current > result {
                result = current
            }
        }
    }

    return result
}

func maxSumArrayWithK(list []int, k int) int {
    result, sum := math.MinInt64, 0

    for i := 0; i < len(list); i++ {
        sum += list[i]
        result = max(result, sum)

        if result == k {
            return k
        }

        if sum < 0 {
            sum = 0
        }
    }

    if result < k {
        return result
    }

    result = math.MinInt64

    for i := 0; i < len(list); i++ {
        sum = 0

        for j := i; j < len(list); j++ {
            sum += list[j]

            if sum == k {
                return k
            }

            if sum > result && sum < k {
                result = sum
            }
        }
    }

    return result
}
