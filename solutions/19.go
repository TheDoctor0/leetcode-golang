package solutions

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}

    for n > 0 {
        head = head.Next
        n--
    }

    prev := dummy

    for head != nil {
        head, prev = head.Next, prev.Next
    }

    prev.Next = prev.Next.Next

    return dummy.Next
}
