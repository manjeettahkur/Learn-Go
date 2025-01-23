package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(findSubstringStd("Manjeet", "jee"))
}

func findSubstringStd(str, sub string) int {
	return strings.Index(str, sub)
}
func reverseWord(input string) string {
	char := []rune(input)
	for i, j := 0, len(char)-1; i < len(char)/2; i, j = i+1, j-1 {
		char[i], char[j] = char[j], char[i]
	}
	return string(char)
}

func magicLatin(str string) string {

	words := strings.Fields(str)
	for i, word := range words {
		if len(word) > 0 && unicode.IsLetter((rune(word[0]))) {
			words[i] = word[1:] + string(word[0]) + "ay"
		}
	}
	return strings.Join(words, " ")

}

func validParentheses(input string) bool {
	stack := []rune{}

	for _, char := range input {
		if char == '(' {
			stack = append(stack, char)
		} else if char == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// )(())(
func validParenthesesTwo(input string) bool {
	leftBracket := 0
	valid := true

	for _, char := range input {
		if char == '(' {
			leftBracket++
		} else if char == ')' {
			leftBracket--
		}
		if leftBracket < 0 {
			valid = false
		}
	}

	if leftBracket == 0 {
		return valid
	}

	return false
}

func nextBigger(n int) int {
	digits := []byte(strconv.Itoa(n))
	pivot := -1
	fmt.Println(digits)

	for i := len(digits) - 2; i >= 0; i-- {
		fmt.Println(digits[i])
		if digits[i] < digits[i+1] {
			pivot = i
			break
		}
	}

	if pivot == -1 {
		return -1
	}

	return n
}

func splitToDigits(n int) []int {
	var digits []int
	digits = []int{20}
	for n > 0 {
		fmt.Println("n%10", n%10)
		fmt.Println("n/=10", n/10)
		digits = append([]int{n % 10}, digits...)
		fmt.Println(digits)
		n /= 10
	}

	return digits
}
