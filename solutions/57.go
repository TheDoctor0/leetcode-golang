package solutions

func insert(intervals [][]int, newInterval []int) [][]int {
    var left, right [][]int

    for _, value := range intervals {
        if value[1] < newInterval[0] {
            left =  append(left, value)
        } else if value[0] > newInterval[1] {
            right = append(right, value)
        } else {
            newInterval[0] = min(newInterval[0], value[0])
            newInterval[1] = max(newInterval[1], value[1])
        }
    }

    return append(append(left, newInterval), right...)
}
