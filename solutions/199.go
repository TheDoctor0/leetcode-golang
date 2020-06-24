package solutions

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    var result = make([]int, 0)

    viewRightSide(&result, root, 0)

    return result
}

func viewRightSide(result *[]int, root *TreeNode, depth int) {
    if root == nil {
        return
    }

    if len(*result) <= depth {
        *result = append(*result, root.Val)
    }

    viewRightSide(result, root.Right, depth + 1)
    viewRightSide(result, root.Left, depth + 1)
}
