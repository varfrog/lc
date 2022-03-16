package palindrome5

func longestPalindrome(s string) string {
	rs := []rune(s)
	strLen := len(rs)
	for winLen := strLen; winLen > 1; winLen-- {
		startPos := 0
		endPos := startPos + winLen
		for {
			if endPos > strLen {
				break
			}
			if isPalindrome(rs[startPos:endPos]) {
				return string(rs[startPos:endPos])
			}

			startPos++
			endPos = startPos + winLen
		}
	}

	return string(rs[0])
}

func isPalindrome(rs []rune) bool {
	i := 0
	j := len(rs) - 1
	for i < j {
		if rs[i] != rs[j] {
			return false
		}
		i++
		j--
	}
	return true
}
