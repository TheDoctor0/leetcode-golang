package solutions

func mergeKLists(lists []*ListNode) *ListNode {
    length := len(lists)

    if length == 0 {
        return nil
    }

    if length == 1 {
        return lists[0]
    }

    l1 := mergeKLists(lists[:length / 2])
    l2 := mergeKLists(lists[length / 2:])

    return mergeLists(l1, l2)
}

func mergeLists(l1, l2 *ListNode) *ListNode {
    var dummy ListNode
    current := &dummy

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            current.Next = l1

            l1 = l1.Next
        } else {
            current.Next = l2

            l2 = l2.Next
        }

        current = current.Next
    }

    if l1 != nil {
        current.Next = l1
    } else if l2 != nil {
        current.Next = l2
    }

    return dummy.Next
}
