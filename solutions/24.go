package solutions

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    dummy := head.Next

    var temp *ListNode
    var previous *ListNode

    for current := head; current != nil && current.Next != nil; current = current.Next.Next {
        temp = current.Next
        current.Next = current.Next.Next
        temp.Next = current
        current = temp

        if previous != nil {
            previous.Next = current
        }

        previous = current.Next
    }

    return dummy
}
