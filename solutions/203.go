package solutions

func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{}
    head, dummy.Next = dummy, head

    for dummy.Next != nil {
        if dummy.Next.Val == val {
            dummy.Next = dummy.Next.Next
        } else {
            dummy = dummy.Next
        }
    }

    return head.Next
}
