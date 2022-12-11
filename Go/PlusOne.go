package main

import "fmt"

func main() {

	TestPlusOne()

}

func plusOne(digits []int) []int {
	i := len(digits) - 1

	digits[i]++
	if digits[i] != 0 && digits[i] != 10 {
		return digits
	}

	for digits[i] == 10 {
		digits[i] = 0
		i--
		// if runs out of digits
		if i < 0 {
			digits = append([]int{1}, digits...)
			break
		}
		digits[i]++
	}
	return digits

}

func TestPlusOne() {

	assert := func(a, b []int) {
		if len(a) != len(b) {
			fmt.Printf("FAIL : GOT %v , EXPECTED %v\n", a, b)
			return
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				fmt.Printf("FAIL : GOT %v , EXPECTED %v\n", a, b)
				return
			}
		}
		fmt.Println("OK")
	}

	assert([]int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	assert([]int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
	assert([]int{1, 0}, plusOne([]int{9}))
	assert([]int{1, 0, 0, 0}, plusOne([]int{9, 9, 9}))
	assert([]int{1, 0, 0, 0, 0}, plusOne([]int{9, 9, 9, 9}))

}
