package solutions

func wiggleMaxLength(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }

    current, peak := nums[0], false

    for i := 1; i < len(nums); i++ {
        if nums[i] != current {
            peak = nums[i] < current

            break
        }
    }

    result := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] == current {
            continue
        }

        if nums[i] < current {
            if peak {
                peak = !peak
                result++
            }
            current = nums[i]
        } else {
            if !peak {
                peak = !peak
                result++
            }

            current = nums[i]
        }
    }

    return result
}
