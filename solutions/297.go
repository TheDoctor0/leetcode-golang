package solutions

import (
    "bytes"
    "fmt"
    "strconv"
    "strings"
)

type Codec struct {}

func Constructor() Codec {
    return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }

    var buffer bytes.Buffer
    queue, size := []*TreeNode{root}, 1

    for size != 0 {
        pop := queue[0]
        queue = queue[1:]

        if pop == nil {
            buffer.WriteString("#!")
            continue
        }

        size--
        buffer.WriteString(fmt.Sprintf("%d!", pop.Val))

        if pop.Left != nil {
            size++
        }
        if pop.Right != nil {
            size++
        }

        queue = append(queue, pop.Left, pop.Right)

    }

    return buffer.String()
}

func (this *Codec) deserialize(data string) *TreeNode {
    nodes := strings.Split(data, "!")

    if len(nodes) == 1 {
        return nil
    }

    value, _ := strconv.Atoi(nodes[0])
    root := &TreeNode{Val: value}

    queue := []*TreeNode{root}
    flags := []int{0}
    index := 0

    for i := 1; i < len(nodes)-1; i++ {
        var node *TreeNode

        if value, errors := strconv.Atoi(nodes[i]); errors == nil {
            node = &TreeNode{Val: value}
            queue = append(queue, node)
            flags = append(flags, 0)
        }

        indexNode := queue[index]

        if flags[index] == 0 {
            indexNode.Left = node
            flags[index] = 1
        } else {
            indexNode.Right = node
            index++
        }
    }

    return root
}
