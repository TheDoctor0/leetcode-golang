package solutions

func levelOrder(root *Node) [][]int {
    var result [][]int

    traverseTree(root, 0, &result)

    return result
}

func traverseTree(root *Node, level int, result *[][]int) {
    if root == nil {
        return
    }

    for _, value := range root.Children {
        traverseTree(value, level + 1, result)
    }

    for len(*result) <= level {
        *result = append(*result, []int{})
    }

    (*result)[level] = append((*result)[level], root.Val)
}
