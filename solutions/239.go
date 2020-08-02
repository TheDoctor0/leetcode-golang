package solutions

type Deque struct {
    indexes []int
}

func (deque *Deque) push(i int) {
    deque.indexes = append(deque.indexes, i)
}

func (deque *Deque) getFirst() int {
    return deque.indexes[0]
}

func (deque *Deque) popFirst() {
    deque.indexes = deque.indexes[1:]
}

func (deque *Deque) getLast() int {
    return deque.indexes[len(deque.indexes) - 1]
}

func (deque *Deque) popLast() {
    deque.indexes = deque.indexes[:len(deque.indexes) - 1]
}

func (deque *Deque) empty() bool {
    return len(deque.indexes) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) < k || k == 0 {
        return make([]int, 0)
    } else if 1 == k {
        return nums
    }

    result := make([]int, len(nums) - k + 1)
    deque := &Deque{}

    for i := range nums {
        if !deque.empty() && (i - k == deque.getFirst()) {
            deque.popFirst()
        }

        for !deque.empty() && nums[deque.getLast()] < nums[i] {
            deque.popLast()
        }

        deque.push(i)

        if i >= k - 1 {
            result[i - k + 1] = nums[deque.getFirst()]
        }
    }

    return result
}