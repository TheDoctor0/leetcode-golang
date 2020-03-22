package solutions

func deleteDuplicates(head *ListNode) *ListNode {
    node := &head

    for *node != nil && (*node).Next != nil {
        value := (*node).Val

        if (*node).Next.Val != value {
            node = &((*node).Next)

            continue
        }

        for *node != nil && (*node).Val == value {
            *node = (*node).Next
        }
    }

    return head
}