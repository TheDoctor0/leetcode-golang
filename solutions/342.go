package solutions

import (
    "math"
)

func isPowerOfFour(num int) bool {
    if num <= 0 {
        return false
    }

    a := math.Log(float64(num)) / math.Log(4)

    return math.Floor(a) == a
}
