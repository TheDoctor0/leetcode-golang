package solutions

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }

    if root.Left != nil {
        if root.Right != nil {
            root.Left.Next = root.Right
        } else {
            root.Left.Next = findNext(root)
        }
    }

    if root.Right != nil {
        root.Right.Next = findNext(root)
    }

    connect(root.Right)
    connect(root.Left)

    return root
}

func findNext(root *Node) *Node {
    for ; root != nil; root = root.Next {
        if root.Next != nil {
            if root.Next.Left != nil {
                return root.Next.Left
            } else if root.Next.Right != nil {
                return root.Next.Right
            }
        }
    }

    return nil
}
