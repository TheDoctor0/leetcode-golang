package solutions

func iterateGrid(xStart, xEnd, yStart, yEnd int, grid *[][]int) *Node{
    nodes := xEnd - xStart

    if nodes == 0 {
        return nil
    }

    if testLeaf(xStart, xEnd, yStart, yEnd, grid) {
        return &Node{(*grid)[xStart][yStart] == 1, true, nil, nil, nil, nil}
    } else {
        return &Node{
            true,
            false,
            iterateGrid(xStart, xStart + nodes / 2, yStart, yStart + nodes / 2, grid),
            iterateGrid(xStart, xStart + nodes / 2, yStart + nodes / 2, yEnd, grid),
            iterateGrid(xStart + nodes / 2, xEnd, yStart, yStart + nodes/ 2, grid),
            iterateGrid(xStart + nodes/ 2, xEnd, yStart + nodes/ 2, yEnd, grid)
        }
    }
}

func testLeaf(xStart, xEnd, yStart, yEnd int, grid *[][]int) bool {
    value := (*grid)[xStart][yStart]

    for i := xStart; i < xEnd; i ++ {
        for j := yStart; j < yEnd; j ++{
            if (*grid)[i][j] != value {
                return false
            }
        }
    }

    return true
}

func construct(grid [][]int) *Node {
    return iterateGrid(0, len(grid), 0, len(grid), &grid)
}
