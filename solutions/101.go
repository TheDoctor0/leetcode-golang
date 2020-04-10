package solutions

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return isSymetricRecursive(root.Left, root.Right)
}

func isSymetricRecursive(p, q *TreeNode) bool {
    if p == nil || q == nil {
        return p == q
    }

    if p.Val != q.Val {
        return false
    }

    return isSymetricRecursive(p.Left, q.Right) && isSymetricRecursive(p.Right, q.Left)
}
