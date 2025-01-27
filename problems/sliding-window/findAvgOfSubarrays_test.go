package slidingwindow

import (
	"reflect"
	"testing"
)

func TestFindAvgOfSubarrays(t *testing.T) {
	tests := []struct {
		arr      []int
		k        int
		expected []float64
	}{
		{
			arr:      []int{1, 3, 2, 6, -1, 4, 1, 8, 2},
			k:        5,
			expected: []float64{2.2, 2.8, 2.4, 3.6, 2.8},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []float64{1.5, 2.5, 3.5, 4.5},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []float64{2.0, 3.0, 4.0},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []float64{3.0},
		},
	}

	for _, test := range tests {
		result := FindAvgOfSubarrays(test.arr, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For array %v and k=%d, expected %v but got %v", test.arr, test.k, test.expected, result)
		}
	}
}

func BenchmarkFindAvgOfSubarrays(b *testing.B) {
    // Test data
    arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
    k := 5

    // Reset timer before the loop
    b.ResetTimer()
    
    // Run the function b.N times
    for i := 0; i < b.N; i++ {
        FindAvgOfSubarrays(arr, k)
    }
}


func BenchmarkFindAvgOfSubarraysImp(b *testing.B) {
    // Test data
    arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
    k := 5

    // Reset timer before the loop
    b.ResetTimer()
    
    // Run the function b.N times
    for i := 0; i < b.N; i++ {
        FindAvgOfSubarrays(arr, k)
    }
}
// Additional benchmark with larger input
func BenchmarkFindAvgOfSubarraysLarge(b *testing.B) {
    // Create larger test data
    arr := make([]int, 1000)
    for i := range arr {
        arr[i] = i
    }
    k := 100

    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        FindAvgOfSubarrays(arr, k)
    }
}

func BenchmarkFindAvgOfSubarraysLargeImp(b *testing.B) {
    // Create larger test data
    arr := make([]int, 1000)
    for i := range arr {
        arr[i] = i
    }
    k := 100

    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        FindAvgOfSubarraysImp(arr, k)
    }
}