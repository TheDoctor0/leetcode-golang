package solutions

func postorderTraversal(root *TreeNode) []int {
    var result []int

    postorderTraverse(root, &result)

    return result
}

func postorderTraverse(root *TreeNode,output *[]int) {
    if root != nil {
        postorderTraverse(root.Left, output)
        postorderTraverse(root.Right, output)

        *output = append(*output, root.Val)
    }
}
