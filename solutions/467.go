package solutions

import "math"

func findSubstringInWraproundString(p string) int {
	if len(p) == 0 {
		return 0
	}

	counter := make([]int, 26)
	start, previous, count := p[0], p[0], 1

	for i := 1; i < len(p); i++ {
		if int(p[i]) - int(previous) == 1 || int(p[i]) - int(previous) == -25 {
			previous = p[i]
			count++
		} else {
			counter[start - 'a'] = maxValue(count, counter[start - 'a'])
			start, previous, count = p[i], p[i], 1
		}
	}

	counter[start - 'a'] = maxValue(count, counter[start - 'a'])
	amount, maxCount, maxCountIndex := 0, 0, 0

	for i, value := range counter {
		if value > maxCount {
			maxCount = value
			maxCountIndex = i
		}
	}

	for i := maxCountIndex + 26; i < maxCountIndex + 52; i++ {
		counter[i % 26] = maxValue(counter[i % 26], counter[(i - 1) % 26]-1)
	}

	for _, value := range counter {
		amount += value
	}

	return amount
}

func maxValue(values ...int) int {
	maxValue := math.MinInt32

	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}
