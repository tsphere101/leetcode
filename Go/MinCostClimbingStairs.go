package main

import "fmt"

func main() {
	cost := []int{10, 15, 20}

	assert(minCostClimbingStairs(cost), 15)
	assert(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}), 6)

}

func minCostClimbingStairs(cost []int) int {

	// take first index

}

func min(arr ...int) int {
	min_ := arr[0]
	for i := range arr {
		if arr[i] < min_ {
			min_ = arr[i]
		}
	}
	return min_
}

func assert(a, b any) {
	if a != b {
		e := fmt.Errorf("%v and %v is not equal", a, b)
		fmt.Println(e)
	}
}
