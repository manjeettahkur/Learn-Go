// Given a string with lowercase letters only, if you are allowed to replace no more than K letters with any letter, find the length of the longest substring having the same letters after replacement.

package slidingwindow

// [  b b], 2 (b) => [a a b b b b b]
func LengthOfLongestSubstring(input string, k int) int {

	left, maxCount, maxLength := 0, 0, 0

	charFre := make([]int, 26)

	for right:=0; right < len(input); right++{
		currentChar := input[right] - 'A'
		charFre[currentChar]++

		maxCount = max(maxCount, charFre[currentChar])

		if (right-left+1) - maxCount > k {
			leftChar := input[left] - 'A'
			charFre[leftChar]--
			left++
		}
		currentLen := right-left+1

		maxLength = max(maxLength, currentLen)

	}

	return maxLength
}
// what to do next 
