package solutions

func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    tail := head
    current := head.Next
    head.Next = nil

    for current != nil {
        node := current
        current = current.Next
        node.Next = nil

        if node.Val > tail.Val {
            tail.Next = node
            tail = node

            continue
        }

        if node.Val < head.Val {
            node.Next = head
            head = node

            continue
        }

        previous := head

        for previous.Next != nil && previous.Next.Val < node.Val {
            previous = previous.Next
        }

        node.Next = previous.Next
        previous.Next = node
    }

    return head
}
