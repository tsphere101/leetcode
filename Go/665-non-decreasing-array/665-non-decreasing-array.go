package main

import "fmt"

func main() {

	cases := [][]int{
		{4, 2, 3},     // true
		{4, 2, 1},     // false
		{3, 4, 2, 3},  // false
		{1, 2, 2, 3},  // true
		{5, 7, 1, 8},  // true
		{5, 7, 1},     // true
		{-1, 4, 2, 3}, // true
	}

	for _, c := range cases {
		fmt.Println(checkPossibilityOptimized(c))
	}
}

func checkPossibilityOptimized(nums []int) bool {
	if len(nums) < 3 {
		return true
	}

	violation := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			violation++

			// if this is the first violation
			if violation == 1 {
				// check out of bound
				if i == 0 {
					// change i to next be value
					nums[i] = nums[i+1]
					continue
				}
				if i+1 == len(nums)-1 {
					// change next value to be i
					nums[i+1] = nums[i]
					continue
				}

				if nums[i+1] < nums[i-1] {
					// change next value to be i
					nums[i+1] = nums[i]
				} else {
					// change i to next be value
					nums[i] = nums[i+1]
				}
			}

			// if this is the second violation
			return false

		}

	}
	return true
}

func checkPossibility(nums []int) bool {
	if len(nums) < 3 {
		return true
	}

	var i, j, k int
	for ; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			j = i
			break
		}
	}

	// finish, no violation
	if i == len(nums)-1 {
		return true
	}

	i++

	k = i

	for ; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			return false
		}
	}

	// j is the first and the only element
	if j-1 == -1 {
		return true
	}

	// k is the last and the only element
	if k+1 >= len(nums) {
		return true
	}

	if nums[k] >= nums[j-1] {
		return true
	}
	if nums[k+1] >= nums[j] {
		return true
	}

	// for this such case [3,4,2,3]
	return false

}
