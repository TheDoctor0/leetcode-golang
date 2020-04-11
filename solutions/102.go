package solutions

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    var result [][]int
    var roots []*TreeNode

    roots = append(roots, root)

    for len(roots) > 0 {
        var currentResult []int
        rootsCount := len(roots)

        for i := 0; i < rootsCount; i++ {
            if roots[i] != nil {
                if roots[i].Left != nil {
                    roots = append(roots, roots[i].Left)
                }

                if roots[i].Right != nil {
                    roots = append(roots, roots[i].Right)
                }

                currentResult = append(currentResult, roots[i].Val)
            }
        }

        roots = roots[rootsCount:]
        result = append(result, currentResult)
    }

    return result
}
