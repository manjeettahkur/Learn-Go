// Given two strings s and p, return an array of all the start indices of p's
// anagrams
//  in s. You may return the answer in any order.

// Example 1:

// Input: s = "cbaebabacd", p = "abc"
// Output: [0,6]
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".

package slidingwindow

func FindStringAnagrams(str, pattern string) []int {

	n := len(pattern)
	matched := 0
	var result []int

	left := 0
	charFreq := make(map[byte]int, n)

	for i := 0; i < n; i++ {
		charFreq[pattern[i]]++
	}

	for right := 0; right < len(str); right++ {

		currentChar := str[right]
		if _, ok := charFreq[currentChar]; ok {

			if charFreq[currentChar] == 0 {
				matched++
			}

			charFreq[currentChar]--

		}

		if matched == len(charFreq) {
			result = append(result, left)
		}

		if right >= n-1 {

			leftChar := str[left]
			left++

			if _, ok := charFreq[leftChar]; ok {
				if charFreq[currentChar] == 0 {
					matched--
				}
				charFreq[leftChar]++
			}

		}

	}

	return result

}
