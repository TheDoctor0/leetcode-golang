package solutions

func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    var slow, fast = head, head

    for {
        if fast == nil || fast.Next == nil {
            return nil
        }

        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            break
        }
    }

    slow = head

    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }

    return fast
}
