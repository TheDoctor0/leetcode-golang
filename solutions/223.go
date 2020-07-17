package solutions

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
    sum1 := (C - A) * (D - B)
    sum2 := (H - F) * (G - E)

    if E >= C || A >= G || F >= D || B >= H {
        return sum1 + sum2
    }

    intersect := (min(C, G) - max(A, E)) * (min(D, H) - max(B, F))

    return sum1 + sum2 - intersect
}
