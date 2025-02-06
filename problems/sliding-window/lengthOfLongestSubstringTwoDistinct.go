// Given a string, find the length of the longest substring in it with at most two distinct characters.

// eceba

package slidingwindow

func LengthOfLongestSubstringTwoDistinct(input string) int {
	str := []rune(input)
	left, maxCount := 0, 0

	charFre := make(map[rune]int, len(str))

	for right, char := range str {

		charFre[char]++

		for len(charFre) > 2 {
			leftChar := str[left]
			charFre[leftChar]--
			if  charFre[leftChar] == 0 {
				delete(charFre, leftChar)
			}

			left++

		}

		if currentCount := right - left + 1; maxCount < currentCount {
			maxCount = currentCount
		}
	}

	return maxCount
}
