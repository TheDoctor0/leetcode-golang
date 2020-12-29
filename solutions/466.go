package solutions

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	i1, i2 := 0, 0
	p1, p2 := 0, 0

	for i1 < n1 {
		if s1[p1] == s2[p2]  {
			p1++
			p2++
		} else {
			p1++
		}

		if p1 >= len(s1) {
			i1++
			p1 = 0
		}

		if p2 >= len(s2) {
			i2++
			p2 = 0
		}
	}

	return i2 / n2
}
