package solutions

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    fast, slow := head, head

    for fast != nil {
        fast = fast.Next

        if fast != nil {
            fast = fast.Next
        }

        slow = slow.Next

        if slow == fast {
            return true
        }
    }

    return false
}