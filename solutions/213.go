package solutions

func rob(nums []int) int {
    if len(nums) == 0{
        return 0
    }

    if len(nums) == 1 {
        return nums[0]
    }

    include, exclude, temp := 0,0,0

    for i := 0; i < len(nums) - 1; i++ {
        temp = include
        include = max(include, exclude + nums[i])
        exclude = temp
    }

    currentMax := include
    include, exclude, temp = 0,0,0

    for i := 1; i < len(nums); i++ {
        temp = include
        include = max(include, exclude + nums[i])
        exclude = temp
    }

    return max(currentMax, include)
}
