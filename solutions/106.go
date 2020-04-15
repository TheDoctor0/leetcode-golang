package solutions

func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }

    value := postorder[len(postorder) - 1]
    root := &TreeNode{
        Val: value,
    }

    var item int
    for item = 0; item < len(inorder); item++ {
        if inorder[item] == value {
            break
        }
    }

    root.Left = buildTree(inorder[:item], postorder[:item])
    root.Right = buildTree(inorder[item + 1:], postorder[item: len(postorder) - 1])

    return root
}
