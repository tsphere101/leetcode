package main

import "fmt"

func main() {

	// case1 := []int{2, 1, 0, 0, 1, 2}
	// case1 := []int{2, 0, 2, 1, 1, 0}
	// case1 := []int{2, 2}
	// case1 := []int{}
	// case1 := []int{2}
	// case1 := []int{1}
	// case1 := []int{0}
	case1 := []int{1, 2, 0}

	sortColors(case1)
	fmt.Println(case1)

}

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	// using modification of quicksort

	red := -1
	blue := len(nums)

	for i := 0; i < len(nums); i++ {
		// move red and blue pointer
		if red+1 >= len(nums) {
			break
		}
		for nums[red+1] == 0 {
			red++
			if red+1 >= len(nums) {
				break
			}
		}

		if blue-1 <= 0 {
			break
		}
		for nums[blue-1] == 2 {
			blue--
			if blue-1 <= 0 {
				break
			}
		}

		// move i pointer
		if i <= red {
			i = red + 1
			if i >= len(nums) {
				break
			}
		}

		if i >= blue {
			// finished sorting
			break
		}

		// if it is red, move red pointer to the right, and swap with i
		if nums[i] < 1 {
			red++
			nums[red], nums[i] = nums[i], nums[red]
			i--
			continue
		}
		// if it is blue, move blue pointer to the left, and swap with i
		if nums[i] > 1 {
			blue--
			nums[blue], nums[i] = nums[i], nums[blue]
			i--
			continue
		}
	}

}
