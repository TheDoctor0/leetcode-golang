package solutions

func rob(root *TreeNode) int {
    result := searchHouses(root)

    return max(result[0], result[1])
}

func searchHouses(root *TreeNode) [2]int {
    if root == nil {
        return [2]int{0, 0}
    }

    left, right := searchHouses(root.Left), searchHouses(root.Right)

    result := [2]int{0, 0}
    result[0] = max(left[0], left[1]) + max(right[0], right[1])
    result[1] = root.Val + left[0] + right[0]

    return result
}
