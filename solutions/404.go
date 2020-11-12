package solutions

func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }

    value := 0

    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
        value = root.Left.Val
    }

    return value + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
