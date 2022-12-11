package main

import "fmt"

func main() {
	a := []int{1, 1, 2}
	b := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(a))
	fmt.Println(removeDuplicates(b))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	current := 0
	for lead := 0; lead < len(nums); lead++ {

		if nums[current] != nums[lead] {
			current++
			nums[current] = nums[lead]
		}
	}
	return current + 1

}
