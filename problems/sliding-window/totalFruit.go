// Given an array of characters where each character represents a fruit tree, you are given two baskets, and your goal is to put the maximum number of fruits in each basket. The only restriction is that each basket can have only one type of fruit.

// You can start with any tree, but you cant skip a tree once you have started. You will pick one fruit from each tree until you cannot, i.e., you will stop when you have to pick from a third fruit type.

// Write a function to return the maximum number of fruits in both baskets.

package slidingwindow

func TotalFruit(fruits []int) int {

	left, maxLength := 0, 0
	count := make(map[int]int, len(fruits))

	for right, v := range fruits {

		count[v]++

		for len(count) > 2 {
			leftVal := fruits[left]
			count[leftVal]--
			if count[leftVal] == 0 {
				delete(count, leftVal)
			}
			left++

		}

		if currentLen := right - left + 1; maxLength < currentLen {
			maxLength = currentLen
		}

	}

	return maxLength
}
