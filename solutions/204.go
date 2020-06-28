package solutions

func countPrimes(n int) int {
    if n < 3 {
        return 0
    }

    count := n / 2
    composite := make([]bool, n)

    for i := 3; i * i < n; i += 2 {
        if composite[i] {
            continue
        }

        for j := i * i; j < n; j += 2 * i {
            if !composite[j] {
                composite[j] = true
                count--
            }
        }
    }

    return count
}
