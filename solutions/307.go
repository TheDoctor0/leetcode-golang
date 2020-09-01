package solutions

type NumArray struct {
    items []int
}

func Constructor(nums []int) NumArray {
    return NumArray{items: nums}
}

func (this *NumArray) Update(i int, value int)  {
    this.items[i] = value
}

func (this *NumArray) SumRange(i int, j int) int {
    subArray, sum := this.items[i: j + 1], 0

    for _, item := range subArray {
        sum += item
    }

    return sum
}
