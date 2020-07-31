package solutions

func deleteNode(node *ListNode) {
    *node = *node.Next
}
