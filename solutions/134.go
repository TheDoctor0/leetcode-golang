package solutions

func canCompleteCircuit(gas []int, cost []int) int {
    totalTank, currentTank, startStation := 0, 0, 0

    for i, value := range gas {
        difference := value - cost[i]

        totalTank += difference
        currentTank += difference

        if currentTank < 0 {
            startStation = i + 1
            currentTank = 0
        }
    }

    if totalTank > -1 {
        return startStation
    }

    return -1
}
