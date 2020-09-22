package solutions

func isSelfCrossing(x []int) bool {
    if len(x) < 4 {
        return false
    }

    for _, value := range x[1 :] {
        if value == 0 {
            return true
        }
    }

    if x[3] >= x[1] && x[2] <= x[0] {
        return true
    }

    spiralIn := x[3] < x[1]

    for i := 4; i < len(x); i ++ {
        if spiralIn {
            if x[i] >= x[i - 2] {
                return true
            }
        } else {
            if x[i] <= x[i - 2] {
                spiralIn = true

                if x[i] + x[i - 4] >= x[i - 2] {
                    x[i - 1] = x[i - 1] - x[i - 3]
                }
            }
        }
    }

    for _, value := range x[1 :] {
        if value == 0 {
            return true
        }
    }

    return false
}
