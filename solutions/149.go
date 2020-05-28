package solutions

import (
    "math/big"
)

func maxPoints(points [][]int) int {
    length := len(points)

    if length == 0 {
        return 0
    }

    result := 1

    for i := 0; i < length - 1; i++ {
        lookup := make(map[string]int, len(points))
        x1, y1 := points[i][0], points[i][1]
        addition, currentMax := 0, 1

        for j := i + 1; j < length; j++ {
            var slopeSymbol string

            x2, y2 := points[j][0], points[j][1]

            if y2 - y1 == 0 && x2 - x1 == 0 {
                addition++

                continue
            } else if x2 - x1 == 0 {
                slopeSymbol = "|"
            } else if y2 - y1 == 0 {
                slopeSymbol = "-"
            } else {
                first := new(big.Rat).SetFloat64(float64(y2 - y1))
                second := new(big.Rat).SetFloat64(float64(x2 - x1))

                slopeSymbol = new(big.Rat).Quo(first, second).FloatString(66)
            }

            if _, ok := lookup[slopeSymbol]; ok {
                lookup[slopeSymbol]++
            } else {
                lookup[slopeSymbol] = 2
            }

            currentMax = max(lookup[slopeSymbol], currentMax)
        }

        currentMax += addition
        result = max(result, currentMax)
    }

    return result
}
