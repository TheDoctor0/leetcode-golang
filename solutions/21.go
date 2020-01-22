package solutions

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := new(ListNode)

    for node := dummy; l1 != nil || l2 != nil; node = node.Next {
        if l1 == nil {
            node.Next = l2

            break
        } else if l2 == nil {
            node.Next = l1

            break
        } else if l1.Val < l2.Val {
            node.Next = l1

            l1 = l1.Next
        } else {
            node.Next = l2

            l2 = l2.Next
        }
    }

    return dummy.Next
}
