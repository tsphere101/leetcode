package main

import "fmt"

func main() {

	// Test cases
	f := getWaysToClimbStairsBetter
	assert(f(0), 0)
	assert(f(1), 1)
	assert(f(2), 2)
	assert(f(3), 3)
	assert(f(4), 5)
	assert(f(5), 8)
	assert(f(6), 13)
	assert(f(7), 21)

}

func assert(a, b any) {
	if a != b {
		fmt.Printf("%v and %v is not equal", a, b)
	}
}

func getWaysToClimbStairsBest(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	a := 1
	b := 2

	for i := 2; i < n+1; i++ {
		a, b = b, a+b
	}
	return a
}

func getWaysToClimbStairsBetter(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	arr := []int{0, 1, 2}

	for i := 3; i < n+1; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}

	return arr[len(arr)-1]

}

func getWaysToClimbStairs(from, to int) int {
	if from > to {
		return 0
	}

	if from == to {
		return 1
	}

	var result int
	result += getWaysToClimbStairs(from+1, to)
	result += getWaysToClimbStairs(from+2, to)
	return result
}
