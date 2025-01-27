package main

import (
	"fmt"
	"github.com/manjeettahkur/learn-go/problems/sliding-window"
)

func main() {

	fmt.Println(slidingwindow.FindAvgOfSubarraysImp([]int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5))

}

// func findSumofMinandMaxSubArray(arr []int, k int) int {

// 	sum := 0
// 	var maxDeque, minDeque []int

// 	for i := 0; i < len(arr); i++ {
// 		// Process maxDeque to maintain elements in decreasing order
// 		for len(maxDeque) > 0 && arr[i] >= arr[maxDeque[len(maxDeque)-1]] {
// 			maxDeque = maxDeque[:len(maxDeque)-1]
// 		}
// 		maxDeque = append(maxDeque, i)

// 		// Remove elements out of the current window from maxDeque
// 		for len(maxDeque) > 0 && maxDeque[0] <= i-k {
// 			maxDeque = maxDeque[1:]
// 		}

// 		// Process minDeque to maintain elements in increasing order
// 		for len(minDeque) > 0 && arr[i] <= arr[minDeque[len(minDeque)-1]] {
// 			minDeque = minDeque[:len(minDeque)-1]
// 		}
// 		minDeque = append(minDeque, i)

// 		// Remove elements out of the current window from minDeque
// 		for len(minDeque) > 0 && minDeque[0] <= i-k {
// 			minDeque = minDeque[1:]
// 		}

// 		// Accumulate sum once the window is formed
// 		if i >= k-1 {
// 			sum += arr[maxDeque[0]] + arr[minDeque[0]]
// 		}
// 	}

// 	return sum
// }

// func findAvgOfSubArray(arr []int, k int) []int {
// 	var res []int

// 	for i := 0; i <= len(arr)-k; i++ {
// 		var sum int
// 		for j := i; j < i+k; j++ {
// 			sum += arr[j]
// 		}
// 		res = append(res, sum)
// 	}

// 	return res
// }

// func FindAvgOfSubArrSW(arr []int, k int) []float64 {
// 	res := make([]float64, 0, len(arr)-k+1)
// 	var sum int
// 	var i int
// 	for j := 0; j < len(arr); j++ {
// 		sum += arr[j]

// 		fmt.Println("j", j, "k-1", k-1)

// 		if j >= k-1 {
// 			res = append(res, float64(sum)/float64(k))
// 			sum -= arr[i]
// 			i++
// 		}

// 	}

// 	return res
// }

// func findSubstringStd(str, sub string) int {
// 	n := len(str)
// 	m := len(sub)

// 	if n == 0 {
// 		return 0
// 	}

// 	if m > n {
// 		return -1
// 	}

// 	for i := 0; i < n-m; i++ {
// 		match := true

// 		for j := 0; j < m; j++ {
// 			if str[i+j] != sub[j] {
// 				match = false
// 				break
// 			}
// 		}

// 		if match {
// 			return i
// 		}
// 	}

// 	return -1

// }

// func reverseWord(input string) string {
// 	char := []rune(input)
// 	for i, j := 0, len(char)-1; i < len(char)/2; i, j = i+1, j-1 {
// 		char[i], char[j] = char[j], char[i]
// 	}
// 	return string(char)
// }

// func magicLatin(str string) string {

// 	words := strings.Fields(str)
// 	for i, word := range words {
// 		if len(word) > 0 && unicode.IsLetter((rune(word[0]))) {
// 			words[i] = word[1:] + string(word[0]) + "ay"
// 		}
// 	}
// 	return strings.Join(words, " ")

// }

// func validParentheses(input string) bool {
// 	stack := []rune{}

// 	for _, char := range input {
// 		if char == '(' {
// 			stack = append(stack, char)
// 		} else if char == ')' {
// 			if len(stack) == 0 || stack[len(stack)-1] != '(' {
// 				return false
// 			}
// 			stack = stack[:len(stack)-1]
// 		}
// 	}

// 	return len(stack) == 0
// }

// // )(())(
// func validParenthesesTwo(input string) bool {
// 	leftBracket := 0
// 	valid := true

// 	for _, char := range input {
// 		if char == '(' {
// 			leftBracket++
// 		} else if char == ')' {
// 			leftBracket--
// 		}
// 		if leftBracket < 0 {
// 			valid = false
// 		}
// 	}

// 	if leftBracket == 0 {
// 		return valid
// 	}

// 	return false
// }

// func nextBigger(n int) int {
// 	digits := []byte(strconv.Itoa(n))
// 	pivot := -1
// 	fmt.Println(digits)

// 	for i := len(digits) - 2; i >= 0; i-- {
// 		fmt.Println(digits[i])
// 		if digits[i] < digits[i+1] {
// 			pivot = i
// 			break
// 		}
// 	}

// 	if pivot == -1 {
// 		return -1
// 	}

// 	return n
// }

// func splitToDigits(n int) []int {
// 	var digits []int
// 	digits = []int{20}
// 	for n > 0 {
// 		fmt.Println("n%10", n%10)
// 		fmt.Println("n/=10", n/10)
// 		digits = append([]int{n % 10}, digits...)
// 		fmt.Println(digits)
// 		n /= 10
// 	}

// 	return digits
// }
