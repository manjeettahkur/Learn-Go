// Given an array of positive numbers and a positive number K, find the maximum sum of any contiguous subarray of size K.

// arr = [2, 1, 5, 1, 3, 2],  k=2

package slidingwindow

func MaxSubarrayOfSizeK(arr []int, k int) int {
	arrLen := len(arr)

	i := 0
	max := 0
	sum := 0
	for j := 0; j < arrLen; j++ {
		sum += arr[j]

		if j >= k-1 {

			if sum > max {
				max = sum
			}
			sum -= arr[i]
			i++

		}

	}

	return max

}
