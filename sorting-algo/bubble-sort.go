// Bubble sort algorithm
// [5,7,2,6,8,1,9,0] = [5,2,7,6,8,1,9,0]
//  i, , j

// [5,2,6,7,1,8,0,9]

package sorting

import "fmt"

func BubbleSort(arr []int) {

	l := len(arr)
	fmt.Println("l", l)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
