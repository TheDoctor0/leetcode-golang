package solutions

type SummaryRanges struct {
    data []int
}

func Constructor() SummaryRanges {
    return SummaryRanges{}
}

func (this *SummaryRanges) bisectLeft(x int) int {
    var left, right = 0, len(this.data)

    for left < right {
        var point = left + (right - left) / 2

        if this.data[point] < x {
            left = point + 1
        } else {
            right = point
        }
    }

    return left
}

func (this *SummaryRanges) bisectRight(x int) int {
    var left, right = 0, len(this.data)

    for left < right {
        var point = left + (right - left) / 2

        if x < this.data[point] {
            right = point
        } else {
            left = point + 1
        }
    }

    return left
}

func (this *SummaryRanges) AddRange(first int, second int)  {
    var left, right = this.bisectLeft(first), this.bisectRight(second)
    var current []int

    current = append(current, this.data[0:left]...)

    if left % 2 == 0 {
        current = append(current, first)
    }

    if right % 2 == 0 {
        current = append(current, second)
    }

    current = append(current, this.data[right:]...)
    this.data = current
}

func (this *SummaryRanges) AddNum(val int)  {
    this.AddRange(val, val+1)
}

func (this *SummaryRanges) GetIntervals() [][]int {
    var result [][]int

    for i := 0; i < len(this.data); i += 2 {
        result = append(result, []int{this.data[i], this.data[i + 1] - 1})
    }

    return result
}
