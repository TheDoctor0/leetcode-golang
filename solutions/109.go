package solutions

func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }

    var previous *ListNode
    slow, fast := head, head

    for fast.Next != nil && fast.Next.Next != nil {
        previous = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    if previous != nil {
        previous.Next = nil
    } else {
        head = nil
    }

    return &TreeNode{
        Val:   slow.Val,
        Left:  sortedListToBST(head),
        Right: sortedListToBST(slow.Next),
    }
}
