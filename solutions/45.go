package solutions

func jump(nums []int) int {
    maxStep := len(nums) - 1

    currentPosition := 0
    steps := 0

    if maxStep == 0 {
        return 0
    }

    for {
        nextMaxStep := nums[currentPosition]
        steps++

        if nextMaxStep >= maxStep {
            break
        }

        max := 0
        nextMove := currentPosition + 1

        for i := currentPosition + 1; i <= currentPosition + nextMaxStep && i < maxStep; i++ {
            if nums[i] >= max || i+nums[i] > max+nextMove {
                max = nums[i]
                nextMove = i
            }
        }

        if nextMove + max >= maxStep {
            steps++

            break
        }

        currentPosition = nextMove

    }

    return steps
}
