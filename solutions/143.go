package solutions

func reorderList(head *ListNode)  {
    if head == nil {
        return
    }

    var reversed *ListNode

    middle := findListMiddle(head)
    next := middle.Next
    middle.Next = nil

    if next != nil {
        reversed = reverseList(next)
    }

    if reversed != nil {
        mergeList(head, reversed)
    }
}

func reverseList(head *ListNode) *ListNode {
    var previous *ListNode

    current, next := head, head.Next

    for ; current != nil ; {
        current.Next = previous
        previous = current
        current = next

        if next != nil {
            next = next.Next
        }
    }

    return previous
}

func findListMiddle(head *ListNode) *ListNode {
    fast, slow := head, head
    moveSlow := false

    for ; fast.Next != nil; {
        fast = fast.Next

        if moveSlow {
            moveSlow = false
            slow = slow.Next
        } else {
            moveSlow = true
        }
    }

    return slow
}

func mergeList(firstHead *ListNode, secondHead *ListNode) {
    for ; secondHead != nil && firstHead != nil; {
        firstNext := firstHead.Next
        secondNext := secondHead.Next

        firstHead.Next = secondHead
        secondHead.Next = firstNext

        firstHead = firstNext
        secondHead = secondNext
    }
}