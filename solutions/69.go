package solutions

func mySqrt(x int) int {
    if x == 0 {
        return 0
    }

    result := 0.5 * float64(x)

    for i := 0 ; i < 20; i++ {
        approximation := 0.5 * (result + float64(x) / result)
        result = approximation
    }

    return int(result)
}
