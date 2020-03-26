package solutions

func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return head
    }

    smaller, greater := &ListNode{0, nil}, &ListNode{0, nil}
    left, right := smaller, greater

    for head != nil {
        if head.Val < x {
            left.Next = head
            left = left.Next
        } else {
            right.Next = head
            right = right.Next
        }

        head = head.Next
    }

    right.Next = nil
    left.Next = greater.Next

    return smaller.Next
}
