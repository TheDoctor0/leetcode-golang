package solutions

type BSTIterator struct {
    stack []*TreeNode
}

type GraphNode struct {
    Val int
    Neighbors []*GraphNode
}

type ListNode struct {
    Val int
    Next *ListNode
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
    Random *Node
}
