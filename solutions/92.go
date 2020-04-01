package solutions

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    node := &ListNode{ Next: head }
    previous := node

    for i := 1; i < m; i++ {
        previous = previous.Next
        head = head.Next
    }

    left := head
    right := head.Next

    for i := m; i < n; i++ {
        next := right.Next

        right.Next = head
        head = right
        right = next
    }

    left.Next = right
    previous.Next = head

    return node.Next
}