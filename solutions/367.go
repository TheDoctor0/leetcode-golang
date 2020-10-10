package solutions

func isPerfectSquare(num int) bool {
    low, high := 1, num

    for low <= high {
        medium := (low + high) >> 1
        current := medium * medium

        if num == current {
            return true
        }

        if current < num {
            low = medium + 1
        } else {
            high = medium - 1
        }
    }

    return false
}
