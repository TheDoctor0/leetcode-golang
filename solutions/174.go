package solutions

func calculateMinimumHP(dungeon [][]int) int {
    rows, columns := len(dungeon) - 1, len(dungeon[0]) - 1

    for row := rows; row >= 0; row-- {
        for column := columns; column >= 0; column-- {
            if row == rows && column == columns {
                dungeon[row][column] = max(1, 1 - dungeon[row][column])

                continue
            }

            if row == rows {
                dungeon[row][column] = max(1, dungeon[row][column + 1] - dungeon[row][column])

                continue
            }

            if column == columns {
                dungeon[row][column] = max(1, dungeon[row + 1][column] - dungeon[row][column])

                continue
            }

            dungeon[row][column] = max(1, min(dungeon[row + 1][column], dungeon[row][column + 1]) - dungeon[row][column])
        }
    }

    return dungeon[0][0]
}
