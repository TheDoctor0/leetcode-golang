package solutions

func searchRange(nums []int, target int) []int {
    left := -1

    for i, j := 0, len(nums); i < j; {
        middle := (i + j) / 2

        if nums[middle] == target {
            left = middle
            j = middle

            continue
        }

        if nums[middle] < target {
            i = middle + 1

            continue
        }

        j = middle
    }

    right := -1

    for i, j := 0, len(nums); i < j; {
        middle := (i + j) / 2

        if nums[middle] == target {
            right = middle
            i = middle + 1

            continue
        }

        if nums[middle] < target {
            i = middle + 1

            continue
        }

        j = middle
    }

    return []int{left, right}
}
