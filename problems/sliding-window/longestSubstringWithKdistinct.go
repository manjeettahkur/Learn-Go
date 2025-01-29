// Given a string, find the length of the longest substring in it with no more than K distinct characters.

// You can assume that K is less than or equal to the length of the given string.

package slidingwindow

// araaci => araa , 2
// [a, r, a, a, c, i]

func LongestSubstringWithKdistinct(input string, k int) int {
	s := []rune(input)
	occr := make(map[rune]int, len(s))
	left, maxLength := 0, 0

	for right, char := range s {

		occr[char]++

		if len(occr) > k {
			leftChar := s[left]
			occr[leftChar]--
			if occr[leftChar] == 0 {
				delete(occr, leftChar)
			}
			leftChar++
		}

		if currentLen := right - left + 1; maxLength < currentLen {
			maxLength = currentLen
		}

	}

	return maxLength
}
