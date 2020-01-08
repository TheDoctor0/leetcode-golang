package solutions

func twoSum(nums []int, target int) []int {
    lookupMap := make(map[int]int)

    for i, value := range nums {
        j, complement := lookupMap[-value]

        if complement {
            return []int{j, i}
        }

        lookupMap[value - target] = i
    }

    return []int{}
}