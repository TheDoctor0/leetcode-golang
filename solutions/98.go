package solutions

func isValidBST(root *TreeNode) bool {
    if root == nil{
        return true
    }

    return validateBST(root, nil, nil)
}

func validateBST(root *TreeNode, min, max interface{}) bool {
    if root == nil {
        return true
    }

    if (min != nil && root.Val <= min.(int)) || (max != nil && root.Val >= max.(int)) {
        return false
    }

    return validateBST(root.Left, min, root.Val) && validateBST(root.Right, root.Val, max)
}
