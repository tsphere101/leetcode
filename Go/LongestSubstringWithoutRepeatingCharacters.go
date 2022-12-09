package main

import (
	"fmt"
)

func main() {

	f := lengthOfLongestSubstringWindowSliding
	assert(f("abcdefghi"), 9)
	assert(f("pwwkew"), 3)
	assert(f("bbbbb"), 1)

}

func assert(a, b any) {
	if a == b {
		fmt.Println("OK")
		return
	}
	fmt.Printf("FAIL : %v is not equal to %v\n", a, b)
}

func lengthOfLongestSubstringWindowSliding(s string) int {
	var left int
	var right int
	var result int
	for ; right < len(s); right++ {

		// while the window contains the repeated char, move the left pointer
		for contains(left, right, s, s[right]) {
			left++
		}

		// update the result, if longest
		if right-left+1 > result {
			result = right - left + 1
		}
	}
	return result
}

func contains(left int, right int, s string, c byte) bool {
	for i := left; i < right; i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	// Length of the longest substring without repeating
	result := 0

	// Loop through s and check
	for i := 0; i < len(s); i++ {
		// substring is the variable for checking the repeating in the substring
		var substring string

		for j := i; j < len(s); j++ {
			// Check if the next character is repeating.
			if in(substring, string(s[j])) {
				break
			}
			substring += string(s[j])
		}
		if len(substring) > result {
			result = len(substring)
		}

		// if the current longest length is more than half of the rest
		if len(s)-i < result {
			break
		}

	}
	return result

}

func in(s string, c string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == c[0] {
			return true
		}
	}
	return false
}
