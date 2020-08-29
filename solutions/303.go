package solutions

type NumArray struct {
    sum []int
}

func Constructor(nums []int) NumArray {
    if len(nums) == 0 {
        return NumArray{}
    }

    sum := make([]int, len(nums) + 1)
    sum[0], sum[1] = 0, nums[0]

    for i := 2; i < len(sum); i++ {
        sum[i] = sum[i - 1] + nums[i - 1]
    }

    return NumArray{sum: sum}
}

func (this *NumArray) SumRange(i int, j int) int {
    return this.sum[j + 1] - this.sum[i]
}
