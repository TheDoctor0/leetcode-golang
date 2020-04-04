package solutions

func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }

    return generateTree(1, n)
}

func generateTree(start, end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }

    var result []*TreeNode

    for s := start; s <= end; s++ {
        left := generateTree(start, s - 1)
        right := generateTree(s + 1, end)

        for _, l := range left {
            for _, r := range right {
                result = append(result, &TreeNode{Val: s, Left: l, Right: r})
            }
        }
    }

    return result
}
