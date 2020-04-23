package solutions

func flatten(root *TreeNode) {
    if root == nil {
        return
    }

    if root.Left != nil {
        flatten(root.Left)

        node := root.Left

        for node.Right != nil {
            node = node.Right
        }

        flatten(root.Right)

        node.Right = root.Right
        root.Right = root.Left
        root.Left = nil

        return
    }

    flatten(root.Right)
}
