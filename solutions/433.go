package solutions

func minMutation(start string, end string, bank []string) int {
    set := make(map[string]struct{})

    for _, value := range bank {
        set[value] = struct{}{}
    }

    queue := []string{start}
    i := 0

    for len(queue) != 0 {
        i++
        length := len(queue)

        for _, firstValue := range queue {
            for secondValue := range set {
                if isNear(firstValue, secondValue) {
                    if secondValue == end {
                        return i
                    }

                    queue = append(queue, secondValue)

                    delete(set, secondValue)
                }
            }
        }

        queue = queue[length:]
    }

    return -1
}

func isNear(firstString string, secondString string) bool {
    mutation := 0

    for i := 0; i < 8; i++ {
        if firstString[i] != secondString[i] {
            mutation++
        }
    }

    return mutation == 1
}