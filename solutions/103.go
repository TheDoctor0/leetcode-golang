package solutions

func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    var result [][]int
    queue := []*TreeNode{root}

    for ascending := true; len(queue) > 0; ascending = !ascending {
        var level []int
        length := len(queue)

        for i := 0; i < length; i++ {
            current := queue[i]

            if ascending {
                level = append(level, current.Val)
            } else {
                level = append(level, queue[length - 1 - i].Val)
            }

            if current.Left != nil {
                queue = append(queue, current.Left)
            }

            if current.Right != nil {
                queue = append(queue, current.Right)
            }
        }

        queue = queue[length:]
        result = append(result, level)
    }

    return result
}