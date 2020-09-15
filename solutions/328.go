package solutions

func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    odd, even, current := head, head.Next, head.Next.Next
    currentOdd, currentEven := odd, even
    isOdd := true

    for current != nil {
        if isOdd {
            currentOdd.Next = current
            currentOdd = currentOdd.Next
        } else {
            currentEven.Next = current
            currentEven = currentEven.Next
        }

        isOdd = !isOdd
        current = current.Next
    }

    currentEven.Next = nil
    currentOdd.Next = even

    return odd
}
