package solutions

func lowestCommonAncestor(node, p, q *TreeNode) *TreeNode {
    var ancestorsP []*TreeNode
    ancestors(node, p.Val, &ancestorsP)

    var ancestorsQ []*TreeNode
    ancestors(node, q.Val, &ancestorsQ)

    result := node

    for i := 0; i < len(ancestorsP) && i < len(ancestorsQ); i++ {
        if ancestorsP[len(ancestorsP) - 1 - i].Val == ancestorsQ[len(ancestorsQ)- 1 - i].Val {
            result = ancestorsP[len(ancestorsP) - 1 - i]
        } else {
            break
        }
    }

    return result
}

func ancestors(node *TreeNode, value int, nodes *[]*TreeNode) bool {
    if node == nil {
        return false
    }

    if node.Val == value || ancestors(node.Left, value, nodes) || ancestors(node.Right, value, nodes) {
        *nodes = append(*nodes, node)

        return true
    }

    return false
}
