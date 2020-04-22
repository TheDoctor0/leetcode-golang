package solutions

func pathSum(root *TreeNode, sum int) [][]int {
    result := make([][]int, 0)
    current := make([]int, 0)

    backtrackPathSum(&result, &current, root, sum)

    return result
}

func backtrackPathSum(result *[][]int, current *[]int, root *TreeNode, sum int) {
    if root == nil {
        return
    }

    sum -= root.Val
    *current = append(*current, root.Val)

    if root.Left == nil && root.Right == nil && sum == 0 {
        temp := make([]int, len(*current))

        copy(temp, *current)

        *result = append(*result, temp)
        *current = (*current)[: len(*current) - 1]

        return
    }

    backtrackPathSum(result, current, root.Left, sum)
    backtrackPathSum(result, current, root.Right, sum)

    *current = (*current)[: len(*current) - 1]
}
