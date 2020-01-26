package solutions

func reverseKGroup(head *ListNode, k int) *ListNode {
    if k == 1 {
        return head
    }

    dummy := head

    for i := 0; i < k; i++ {
        if dummy == nil {
            return head
        }

        dummy = dummy.Next
    }

    previous, current, next := reverseKGroup(dummy, k), head, head.Next

    for current != dummy {
        next, current.Next = current.Next, previous
        previous, current = current, next
    }

    return previous
}
