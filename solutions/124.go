package solutions

import (
    "math"
)

func maxPathSum(root *TreeNode) int {
    result := math.MinInt32

    sumMaxPath(root, &result)

    return result
}

func sumMaxPath(root *TreeNode, result *int) int {
    if root == nil {
        return 0
    }

    left := sumMaxPath(root.Left, result)
    right := sumMaxPath(root.Right, result)

    *result = max(left + right + root.Val, *result)

    return max(0, max(left, right) + root.Val)
}
