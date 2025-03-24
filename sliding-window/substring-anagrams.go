package sliding_window

func substringAnagram(str, t string) int {
	expectedFreq := calculatefrq(t)
	windowSize := len(t)
	currentWindowFrew := [26]int{}
	left, right := 0, 0
	count := 0
	for right < len(str) {
		index := str[right] - 'a'
		currentWindowFrew[index] += 1
		if right-left+1 == windowSize {
			if expectedFreq == currentWindowFrew {
				count++

			}
			leftIndex := str[left] - 'a'
			currentWindowFrew[leftIndex] -= 1
			left++
		}
		right++
	}
	return count
}

func calculatefrq(str string) [26]int {
	freq := [26]int{}
	for _, char := range str {
		index := char - 'a'
		freq[index] += 1
	}
	return freq
}
