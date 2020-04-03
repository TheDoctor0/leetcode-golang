package solutions

func inorderTraversal(root *TreeNode) []int {
    var stack []*TreeNode
    var result []int

    current := root

    for current != nil || len(stack) > 0 {
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }

        current = stack[len(stack) - 1]
        stack = stack[: len(stack) - 1]
        result = append(result, current.Val)
        current = current.Right
    }

    return result
}
