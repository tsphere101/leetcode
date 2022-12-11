package main

import "fmt"

func main() {
	f := mySqrt
	assert(f(4), 2)
	assert(f(8), 2)
	assert(f(9), 3)
	assert(f(2), 1)
	assert(f(1), 1)
	assert(f(0), 0)
	assert(f(1024), 32)

}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left := 0
	right := x
	mid := (right - left) / 2

	for !(mid*mid <= x && (mid+1)*(mid+1) > x) {
		fmt.Println(mid)
		if mid*mid > x {
			right = mid
			mid = mid / 2
		} else {
			left = mid
			mid += (right - mid) / 2
		}
	}
	return mid

}

func found(i, x int) bool {
	if i*i <= x && (i+1)*(i+1) > x {
		return true
	}
	return false
}

func assert(a, b any) {
	if a == b {
		fmt.Println("OK")
		return
	}

	fmt.Printf("GOT %v, EXPECTED %v\n", a, b)
}
