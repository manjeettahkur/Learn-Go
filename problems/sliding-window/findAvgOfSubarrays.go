package slidingwindow

// Given an array, find the average of all contiguous subarrays of size K in it.
// Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
func FindAvgOfSubarrays(arr []int, k int) []float64 {

	var windowStart int
	var sum int
	arrLen := len(arr)
	var result []float64

	for windowEnd := 0; windowEnd < arrLen; windowEnd++ {

		sum += arr[windowEnd]

		if windowEnd >= k-1 {

			result = append(result, float64(sum)/float64(k))
			sum -= arr[windowStart]
			windowStart++
		}
	}

	return result

}

func FindAvgOfSubarraysImp(arr []int, k int) []float64 {

	arrLen := len(arr)
	m := arrLen - k + 1

	if k <= 0 || m <= 0 {
		return []float64{}
	}
	result := make([]float64, m)
	invK := 1.0 / float64(k)
	sum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < arrLen; windowEnd++ {
		sum += arr[windowEnd]
		if windowEnd >= k-1 {
			result[windowStart] = float64(sum) * invK

			sum -= arr[windowStart]
			windowStart++
		}
	}

	return result

}
