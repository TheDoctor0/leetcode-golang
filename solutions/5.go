package solutions

func longestPalindrome(s string) string {
	length := len(s)

	if length <= 1 {
		return s
	}

	var begin, max int

	for i := 0; i < length; {
		left, right := i, i

		for right < length - 1 && s[right] == s[right + 1] {
			right++
		}

		i = right + 1

		for left > 0 && right < length - 1 && s[left - 1] == s[right + 1] {
			left--
			right++
		}

		if max < right - left {
			begin = left
			max = right - left
		}
	}

	return s[begin : begin + max + 1]
}
