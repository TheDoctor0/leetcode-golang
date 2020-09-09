package solutions

func bulbSwitch(n int) int {
    i := 1

    for i * i <= n {
        i++
    }

    return i - 1
}
