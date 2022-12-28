package main

import "fmt"

func main() {

	cases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{5, 4, 3, 2, 1},
	}

	for _, c := range cases {
		fmt.Println(maxArea(c))
	}

}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	ans := 0

	for left <= right {

		ans = max(min(height[left], height[right])*abs(left-right), ans)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
