package palindrome

// makePalindrome helps in making the string palindrome with minimal changes
func makePalindrome(s []rune, left, right, k int) (int, bool) {
	if left >= right {
		return k, true
	}

	if s[left] == s[right] {
		return makePalindrome(s, left+1, right-1, k)
	}

	if k == 0 {
		return k, false
	}

	if s[left] > s[right] {
		s[right] = s[left]
	} else {
		s[left] = s[right]
	}

	return makePalindrome(s, left+1, right-1, k-1)
}

// maximizePalindrome helps in maximizing the palindrome value
func maximizePalindrome(s []rune, left, right, k int) (bool, string) {
	if left > right {
		return true, string(s)
	}

	if s[left] == s[right] {
		remaining, result := maximizePalindrome(s, left+1, right-1, k)
		if !remaining {
			return false, result
		}
		return true, result
	}

	if k > 0 {
		if s[left] != '9' {
			s[left] = '9'
			k--
		}
		if s[right] != '9' {
			s[right] = '9'
			k--
		}
	}

	remaining, result := maximizePalindrome(s, left+1, right-1, k)
	if !remaining {
		return false, result
	}
	return true, result
}

// HighestPalindrome returns the highest value palindrome possible with at most k changes
func HighestPalindrome(s string, k int) string {
	runeStr := []rune(s)
	initialK := k
	newK, ok := makePalindrome(runeStr, 0, len(runeStr)-1, k)

	if !ok {
		return "-1"
	}

	_, result := maximizePalindrome(runeStr, 0, len(runeStr)-1, newK+(initialK-newK)/2)

	return result
}
