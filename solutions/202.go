package solutions

func isHappy(n int) bool {
    set := make(map[int]struct{})

    return isHappyNumber(set, n)
}

func isHappyNumber(set map[int]struct{}, number int) bool {
    if number == 1 {
        return true
    }

    if _, ok := set[number]; ok {
        return false
    } else {
        set[number] = struct{}{}
    }

    remaining := 0

    for {
        if number/ 10 == 0 {
            remaining += (number % 10) * (number % 10)

            break
        } else {
            remaining += (number % 10) * (number % 10)
            number /= 10
        }
    }

    return isHappyNumber(set, remaining)
}
