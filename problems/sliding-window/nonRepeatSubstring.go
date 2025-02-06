package slidingwindow

// abccde
// dvdf
func NonRepeatSubstring(s string) int {

	left, maxCount := 0, 0
	strRune := []rune(s)
	charFre := make(map[rune]int, len(strRune))

	for right, char := range strRune {

		if _, ok := charFre[char]; ok {
			left = max(left, charFre[char]+1)
		}

		charFre[char] = right

		if currCount := right - left + 1; maxCount < currCount {
			maxCount = currCount
		}

	}

	return maxCount

}

func NonRepeatSubstringUsingChar(s string) int {
	if l := len(s); l < 2 {
		return l
	}

	last := [123]int{}
	max, left := 0, 0

	for right, char := range s {

		if last[char] > left{
			left = last[char]
		} else if d := right-left+1; d > max {
			max = d
		}
		last[char] = right + 1
	}

	return max
}
