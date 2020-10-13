package solutions

var modulo = 1337

func superPow(a int, b []int) int {
    if len(b) == 1 {
        return pow(a, b[0])
    }

    result := superPow(a, b[:len(b) - 1])
    result = pow(result, 10)
    result = result * pow(a, b[len(b) - 1]) % modulo

    return result
}

func pow(a int, b int) int {
    if a == 0 {
        return 0
    }

    if b == 0 {
        return 1
    }

    result := pow(a, b >> 1)
    result = result * result % modulo

    if b & 1 == 1 {
        result = result * a % modulo
    }

    return result
}
