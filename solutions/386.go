package solutions

func lexicalOrder(n int) []int {
    result := []int{1}
    number := 1

    for i := 1; i < n; i++ {
        number = next(number, n)
        result = append(result, number)
    }

    return result
}

func next(i int, n int) int {
    if i * 10 <= n {
        return i * 10
    }

    if i + 1 <= n && i % 10 < 9 {
        return i + 1
    }

    if i % 10 == 9  {
        i++
    } else if i + 1 > n {
        i = i / 10 + 1
    } else {
        panic("screwed")
    }

    for i % 10 == 0 {
        i = i / 10
    }

    return i
}
