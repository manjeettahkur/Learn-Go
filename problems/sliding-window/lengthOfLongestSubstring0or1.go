package slidingwindow

// [0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], 2

func LengthOfLongestSubstring0or1(nums []int, k int) int {

	left, maxCountOne, maxLength := 0, 0, 0

	for right, char := range nums {

		if char == 1 {
			maxCountOne++
		}

		if (right-left+1)-maxCountOne > k {
			if nums[left] == 1 {
				maxCountOne--
			}
			left++
		}

		if currentLength := right - left + 1; maxLength < currentLength {
			maxLength = currentLength
		}

	}

	return maxLength
}
