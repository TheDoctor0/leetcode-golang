package solutions

func sortList(head *ListNode) *ListNode {
    if nil == head || nil == head.Next {
        return head
    }

    slow, fast := head, head.Next

    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next

        if fast != nil{
            fast = fast.Next
        }
    }

    newSlow := slow.Next
    slow.Next = nil

    firstHead := sortList(head)
    secondHead := sortList(newSlow)

    return mergeSortedLists(firstHead, secondHead)
}

func mergeSortedLists(first *ListNode, second *ListNode) *ListNode{
    if first == nil {
        return second
    }

    if second == nil {
        return first
    }

    var head *ListNode

    if first.Val < second.Val {
        head = first
        first = first.Next
    } else {
        head = second
        second = second.Next
    }

    var visit = head

    for first != nil && second != nil {
        if first.Val < second.Val {
            visit.Next = first
            visit = first
            first = first.Next
        } else {
            visit.Next = second
            visit = second
            second = second.Next
        }
    }

    if first == nil {
        visit.Next = second
    } else if second == nil {
        visit.Next = first
    }

    return head
}
