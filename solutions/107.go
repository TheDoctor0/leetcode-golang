package solutions

func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    var queue []*TreeNode
    var result [][]int
    queue = append(queue, root)

    for len(queue) > 0 {
        var level []int
        length := len(queue)

        for i := 0; i < length; i++ {
            node := queue[0]
            queue = queue[1:]
            level = append(level, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, level)
    }

    var reversed [][]int

    for i := len(result) - 1; i >= 0; i-- {
        reversed = append(reversed, result[i])
    }

    return reversed
}
