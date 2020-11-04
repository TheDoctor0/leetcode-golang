package solutions

func maxRotateFunction(A []int) int {
    summary, base := 0, 0

    for index, value := range A {
        summary += value
        base += index * value
    }

    max := base

    for i := 0; i < len(A)-1; i++ {
        base += summary
        base -= len(A) * A[len(A) - 1 - i]

        if base > max {
            max = base
        }
    }

    return max
}
