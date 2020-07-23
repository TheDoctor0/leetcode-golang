package solutions

func majorityElement(nums []int) []int {
    firstCount, secondCount, firstCandidate, secondCandidate := 0, 0, 0, 1

    for _, number := range nums {
        if number == firstCandidate {
            firstCount++
        } else if number == secondCandidate {
            secondCount++
        } else if firstCount <= 0 {
            firstCandidate, firstCount = number, 1
        } else if secondCount <= 0 {
            secondCandidate, secondCount = number, 1
        } else {
            firstCount--
            secondCount--
        }
    }

    firstCount, secondCount = 0, 0

    for _, number := range nums {
        if number == firstCandidate {
            firstCount++
        } else if number == secondCandidate {
            secondCount++
        }
    }

    length := len(nums)

    if  firstCount > length / 3 && secondCount > length / 3 {
        return []int{firstCandidate, secondCandidate}
    }

    if firstCount > length / 3 {
        return []int{firstCandidate}
    }

    if secondCount > length / 3 {
        return []int{secondCandidate}
    }

    return []int{}
}
