package solutions

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    if isLeaf(root) {
        return 1
    }

    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func isLeaf(node *TreeNode) bool {
    return node.Left == nil && node.Right == nil
}
