package solutions

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }

    if root.Left != nil {
        root.Left.Next = root.Right

        if root.Next != nil {
            root.Right.Next = root.Next.Left
        }
    }

    connect(root.Left)
    connect(root.Right)

    return root
}
