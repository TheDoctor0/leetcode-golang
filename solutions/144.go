package solutions

func preorderTraversal(root *TreeNode) []int {
    var result []int
    var stack []*TreeNode
    current := root

    for current != nil || len(stack) != 0 {
        for current != nil {
            stack = append(stack, current)
            result = append(result, current.Val)
            current = current.Left
        }

        current = stack[len(stack) - 1]
        stack = stack[: len(stack) - 1]
        current = current.Right
    }

    return result
}
