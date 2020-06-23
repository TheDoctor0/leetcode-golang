package solutions

func rightSideView(root *TreeNode) []int {
    var result = make([]int, 0)

    if root == nil {
        return nil
    }

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
