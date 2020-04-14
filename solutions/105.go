package solutions

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    position := -1

    for i, number := range inorder {
        if preorder[0] == number {
            position = i
        }
    }

    return &TreeNode{
        preorder[0],
        buildTree(preorder[1: position + 1], inorder[0: position]),
        buildTree(preorder[position + 1:], inorder[position + 1:]),
    }
}
