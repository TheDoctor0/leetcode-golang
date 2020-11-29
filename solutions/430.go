package solutions

func flatten(root *Node) *Node {
    if root == nil {
        return root
    }

    dummyHead := &Node{0, nil, root, nil}
    var current, previous *Node = nil, dummyHead

    var stack []*Node
    stack = append(stack, root)

    for len(stack) > 0 {
        current = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]

        previous.Next = current
        current.Prev = previous

        if current.Next != nil {
            stack = append(stack, current.Next)
        }

        if current.Child != nil {
            stack = append(stack, current.Child)
            current.Child = nil
        }

        previous = current
    }

    dummyHead.Next.Prev = nil

    return dummyHead.Next
}
