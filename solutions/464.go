package solutions

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	target, max := desiredTotal, maxChoosableInteger

	if max >= target {
		return true
	}

	if max* (max+ 1) / 2 < target {
		return false
	}

	dp := make([]int, 1 <<max)

	return searchCards(0, max, target, dp)
}

func searchCards(state int, n int, target int, dp []int) bool {
	if dp[state] != 0 {
		return dp[state] == 1
	}

	if target <= 0 {
		return false
	}

	for i := n; i > 0; i-- {
		card := 1 << (i - 1)

		if state & card != 0  {
			continue
		}

		newState := state | card

		if !searchCards(newState, n, target - i, dp) {
			dp[state] = 1

			return true
		}
	}

	dp[state] = -1

	return false
}
