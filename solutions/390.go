package solutions

func lastRemaining(n int) int {
    i, gap, start := 1, 1, 1
    max := n

    for {
        if i % 2 == 1 || n % 2 == 1 {
            start = start + gap

            if start > max {
                return start - gap
            }
        }

        gap = gap << 1
        n = n >> 1
        i++
    }
}
