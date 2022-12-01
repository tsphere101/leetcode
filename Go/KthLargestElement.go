package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	// Create 10 test cases
	for i := 0; i < 100; i++ {
		var nums []int = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100)}
		var k int = rand.Intn(5) + 1
		result := findKthLargest(nums, k)

		if result != recheck(nums, k) {
			fmt.Println("Error: ", nums, k, result, recheck(nums, k))
		} else {
			fmt.Println("OK: ", nums, k, result)
		}

	}

}

func findKthLargest(nums []int, k int) int {
	// Quick select
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, start, end, k int) int {

	// If the array has only one element, return it
	if start == end {
		return nums[start]
	}

	// Randomly select a pivot
	pivot := rand.Intn(end-start) + start

	// Move the pivot to the end of the array
	nums[pivot], nums[end] = nums[end], nums[pivot]

	// Partition the array
	pivot = partition(nums, start, end)

	// If the pivot is the kth largest element, return it
	if pivot == len(nums)-k {
		return nums[pivot]
	}

	// If the pivot is larger than the kth largest element, search the left part of the array
	if pivot > len(nums)-k {
		return quickSelect(nums, start, pivot-1, k)
	}

	// If the pivot is smaller than the kth largest element, search the right part of the array
	return quickSelect(nums, pivot+1, end, k)
}

func partition(nums []int, start, end int) int {
	// Partition the array
	pivot := start
	for i := start; i < end; i++ {
		if nums[i] < nums[end] {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot++
		}
	}
	nums[pivot], nums[end] = nums[end], nums[pivot]
	return pivot
}

func recheck(arr []int, k int) int {
	sort.Ints(arr)
	return arr[len(arr)-k]
}
