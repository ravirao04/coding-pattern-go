package tow_pointer

import "unicode"

func isPalindrome(str string) bool {
	left, right := 0, len(str)-1
	for left < right {
		for left < right && !unicode.IsLetter(rune(str[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(str[right])) {
			right--
		}
		if rune(str[left]) == rune(str[right]) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func calculateFre(str string) [26]int {
	freq := [26]int{}
	for _, char := range str {
		if unicode.IsLetter(char) {
			freq[char-'a'] += 1
		}
	}
	return freq
}
