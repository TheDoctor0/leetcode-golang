package solutions

func nextPermutation(nums []int) {
    if len(nums) == 0 {
        return
    }

    pivot := -1

    for i := len(nums) - 2; i >= 0; i-- {
        if nums[i] < nums[i + 1] {
            pivot = i

            break
        }
    }

    if pivot == -1 {
        reverseNumbers(&nums, 0, len(nums) - 1)

        return
    }

    nextPivot := 0

    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] > nums[pivot] {
            nextPivot = i

            break
        }
    }

    nums[pivot], nums[nextPivot] = nums[nextPivot], nums[pivot]

    reverseNumbers(&nums, pivot + 1, len(nums) - 1)
}

func reverseNumbers(nums *[]int, start, end int) {
    for start < end {
        (*nums)[start], (*nums)[end] = (*nums)[end], (*nums)[start]
        start++
        end--
    }
}
