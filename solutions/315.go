package solutions

func countSmaller(nums []int) []int {
    index := make([]int, 0, len(nums))

    for i := range nums {
        index = append(index, i)
    }

    inversions := make([]int, len(nums), len(nums))

    divideNumbers(0, len(nums) - 1, nums, index, inversions)

    return inversions
}

func divideNumbers(left, right int, nums, index, inversions []int) {
    if left >= right {
        return
    }

    middle := (left + right) / 2

    divideNumbers(left, middle, nums, index, inversions)
    divideNumbers(middle + 1, right, nums, index, inversions)

    mergeNumbers(left, right, nums, index, inversions)
}

func mergeNumbers(left, right int, nums, index, inversions []int) {
    start, end, middle := left, right, (left + right) / 2
    leftMid, rightMid := middle, middle + 1
    smaller, newIndex := 0, make([]int, 0)

    for left <= leftMid && rightMid <= right {
        if nums[index[left]] <= nums[index[rightMid]] {
            newIndex = append(newIndex, index[left])
            inversions[index[left]] += smaller
            left++
        } else {
            newIndex = append(newIndex, index[rightMid])
            rightMid++
            smaller++
        }
    }

    for left <= leftMid {
        inversions[index[left]] += smaller
        newIndex = append(newIndex, index[left])
        left++
    }

    for rightMid <= right {
        newIndex = append(newIndex, index[rightMid])
        rightMid++
    }

    copy(index[start: end + 1], newIndex)
}
