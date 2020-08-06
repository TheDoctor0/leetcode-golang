package solutions

import (
    "strconv"
)

func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return nil
    }

    var result []string

    traverseTreePaths(root, strconv.Itoa(root.Val), &result)

    return result
}

func traverseTreePaths(root *TreeNode, str string, result *[]string) {
    if root.Left == nil && root.Right == nil {
        *result = append(*result, str)

        return
    }

    if root.Left != nil {
        traverseTreePaths(root.Left, str + "->" + strconv.Itoa(root.Left.Val), result)
    }

    if root.Right != nil {
        traverseTreePaths(root.Right, str + "->" + strconv.Itoa(root.Right.Val), result)
    }
}
