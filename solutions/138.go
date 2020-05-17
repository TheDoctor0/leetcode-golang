package solutions

func copyRandomList(head *Node) *Node {
    if head == nil {
        return head
    }

    nodeMap := make(map[*Node]*Node)
    dummy := head

    for ; dummy != nil; dummy = dummy.Next {
        nodeMap[dummy] = &Node{Val: dummy.Val}
    }

    for dummy = head; dummy != nil; dummy = dummy.Next {
        newNode, newNextNode, newRandNode := nodeMap[dummy], nodeMap[dummy.Next], nodeMap[dummy.Random]
        newNode.Next, newNode.Random = newNextNode, newRandNode
    }

    return nodeMap[head]
}
