package solutions

func canMeasureWater(x int, y int, z int) bool {
    if x + y < z {
        return false
    }

    if z == 0 {
        return true
    }

    return z % greatestCommonDivisor(x, y) == 0
}

func greatestCommonDivisor(x int, y int) int {
    if y == 0 {
        return x
    }

    return greatestCommonDivisor(y, x % y)
}
