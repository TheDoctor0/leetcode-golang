package solutions

type GraphNode struct {
    Val int
    Neighbors []*GraphNode
}

func cloneGraph(node *GraphNode) *GraphNode {
    if node == nil {
        return nil
    }

    queue := []*GraphNode{node}
    cloned := map[int]*GraphNode{
        node.Val: {
            Val: node.Val,
            Neighbors: []*GraphNode{},
        },
    }

    for len(queue) > 0 {
        node, queue = queue[0], queue[1:]
        nextNode, _ := cloned[node.Val]

        for _, neighbor := range node.Neighbors {
            if _, visited := cloned[neighbor.Val]; !visited {
                cloned[neighbor.Val] = &GraphNode{
                    Val: neighbor.Val,
                    Neighbors: []*GraphNode{},
                }

                queue = append(queue, neighbor)
            }

            nextNode.Neighbors = append(nextNode.Neighbors, cloned[neighbor.Val])
        }
    }

    return cloned[1]
}
