package solutions

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }

    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }

    return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}
