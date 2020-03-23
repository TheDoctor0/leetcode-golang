package solutions

func deleteDuplicates(head *ListNode) *ListNode {
    node := head

    for head != nil && head.Next != nil {
        if head.Val == head.Next.Val {
            head.Next = head.Next.Next
        } else {
            head = head.Next
        }
    }

    return node
}
