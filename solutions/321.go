package solutions

func maxNumber(nums1 []int, nums2 []int, k int) []int {
    result := make([]int, k)

    for i := 0; i <= k; i++ {
        if len(nums1) >= i && len(nums2) >= k - i {
            first, second := getMax(nums1, i), getMax(nums2, k - i)
            current := mergeMax(first, second)

            if isLarger(current, result) {
                result = current
            }
        }
    }

    return result
}

func getMax(nums []int, k int) []int {
    discard := len(nums) - k
    var stack []int

    for i := 0; i < len(nums); i++ {
        for len(stack) > 0 && discard > 0 && stack[len(stack) - 1] < nums[i] {
            stack = stack[:len(stack) - 1]
            discard--
        }

        stack = append(stack, nums[i])
    }

    return stack[:k]
}

func mergeMax(a []int, b []int) []int {
    var result []int

    for len(a) > 0 && len(b) > 0 {
        if isLarger(a, b) {
            result = append(result, a[0])
            a = a[1:]
        } else {
            result = append(result, b[0])
            b = b[1:]
        }
    }

    if len(a) > 0 {
        result = append(result, a...)
    }

    if len(b) > 0 {
        result = append(result, b...)
    }

    return result
}

func isLarger(a []int, b []int) bool {
    minLength := min(len(a), len(b))

    for i := 0; i < minLength; i++ {
        if a[i] < b[i] {
            return false
        } else if a[i] > b[i] {
            return true
        }
    }

    return len(a) > len(b)
}
