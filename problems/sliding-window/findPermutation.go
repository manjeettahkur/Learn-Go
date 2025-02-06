// Given a string and a pattern, find out if the string contains any permutation of the pattern.
// "eidbaooo" = "ab"
package slidingwindow

func FindPermutation(str, pattern string) bool {

	input := []rune(str)
	left, isMatched := 0, 0

	charFreq := make(map[rune]int, len(pattern))

	for _, char := range pattern {
		charFreq[char]++
	}

	for right, char := range input {

		if _, ok := charFreq[char]; ok {
			charFreq[char]--

			if charFreq[char] == 0 {
				isMatched++
			}

			if len(charFreq) == isMatched {
				return true
			}

		}

		if right >= len(pattern)-1 {

			leftChar := input[left]
			left++

			if _, ok := charFreq[leftChar]; ok{

				if charFreq[leftChar] == 0{
					isMatched--
				}

				charFreq[leftChar]++
			}

			

			if charFreq[leftChar] == 0 {
				isMatched--
			}

			
		}

	}

	return false
}
