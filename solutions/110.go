package solutions

func isBalanced(root *TreeNode) bool {
    return isBalancedTree(root) != -1
}

func isBalancedTree(node *TreeNode) int {
    if node == nil {
        return 0
    }

    left, right := isBalancedTree(node.Left), isBalancedTree(node.Right)

    if left == -1 || right == -1 || abs(left - right) > 1 {
        return -1
    }

    return 1 + max(left, right)
}
