package solutions

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    current := head
    var length = 1

    for current.Next != nil {
        current = current.Next
        length++
    }

    current.Next = head
    j := length - (k % length)

    for i := 1; i <= j; i++ {
        current = current.Next
    }

    head = current.Next
    current.Next = nil

    return head
}
