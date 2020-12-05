package solutions

func pathSum(root *TreeNode, sum int) int {
    return pathSumHelper(root, 0, sum, map[int]int{0: 1})
}

func pathSumHelper(node *TreeNode, currentSum, target int, valueMap map[int]int) int {
    if node == nil {
        return 0
    }

    currentSum += node.Val
    summary := valueMap[currentSum - target]
    valueMap[currentSum]++

    summary += pathSumHelper(node.Left, currentSum, target, valueMap)
    summary += pathSumHelper(node.Right, currentSum, target, valueMap)

    valueMap[currentSum]--

    return summary
}
