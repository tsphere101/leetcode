package main

import "fmt"

func main() {

	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))

}

func coinChange(coins []int, amount int) int {
	// knownCoins holds the number of coin need for the value of its index
	knownCoins := []int{0}

	for value := 1; value < amount+1; value++ {
		var candidates []int
		for _, coin := range coins {
			if value-coin >= 0 {
				candidates = append(candidates, 1+knownCoins[value-coin])
			}
		}
		// remove 0 from candidates
		candidates = removeZero(candidates)

		// no candidates : can't change the coin store -1
		if len(candidates) == 0 {
			knownCoins = append(knownCoins, -1)
			continue
		}

		// multiple candidates, store the minmun value
		knownCoins = append(knownCoins, min(candidates...))

	}

	return knownCoins[amount]

}

func removeZero(arr []int) []int {
	var result []int
	for i := range arr {
		if arr[i] != 0 {
			result = append(result, arr[i])
		}
	}
	return result
}

func min(arr ...int) int {
	rt := arr[0]
	for _, v := range arr {
		if v < rt {
			rt = v
		}
	}
	return rt
}

func coinChangeDFS(coins []int, amount int) int {
	fmt.Println(amount)
	// Base case
	if amount <= 0 {
		return -1
	}

	if len(coins) == 1 {
		if amount%coins[0] == 0 {
			return amount / coins[0]
		}
		return -1
	}

	listOfNumsOfCoins := []int{}
	for _, v := range coins {
		listOfNumsOfCoins = append(listOfNumsOfCoins, coinChangeDFS(coins, amount-v))
	}

	m := minIgnoreNegative(listOfNumsOfCoins...)

	return 1 + m

}

func minIgnoreNegative(arr ...int) int {
	rt := arr[0]
	for _, v := range arr {
		if v < 0 {
			continue
		}
		if v < rt {
			rt = v
		}
	}
	return rt
}

func coinSort(coins []int) {
	// base case
	if len(coins) <= 1 {
		return
	}

	var pivot int
	var bigger int

	for i := 0; i < len(coins); i++ {

		if coins[i] > coins[pivot] {
			bigger++
			coins[i], coins[bigger] = coins[bigger], coins[i]
		}
	}

	coins[pivot], coins[bigger] = coins[bigger], coins[pivot]
	pivot = bigger

	coinSort(coins[:pivot])
	coinSort(coins[pivot+1:])

}

func quickSortDesc(nums []int) {
	// base case
	if len(nums) <= 1 {
		return
	}

	var pivot int
	var bigger int
	var qualifier int

	for ; qualifier < len(nums); qualifier++ {
		if nums[qualifier] > nums[pivot] {
			bigger++
			nums[bigger], nums[qualifier] = nums[qualifier], nums[bigger]
		}
	}
	// swap position of pivot and bigger
	nums[pivot], nums[bigger] = nums[bigger], nums[pivot]
	pivot = bigger

	quickSortDesc(nums[:pivot])
	quickSortDesc(nums[pivot+1:])

}

func quickSort(nums []int) {
	// base cases
	if len(nums) <= 1 {
		return
	}

	var pivot int
	var smaller int
	var qualifier int

	// [pv, sm, sm, sm, b, b, b, q, b, b]

	for ; qualifier < len(nums); qualifier++ {
		// if the qualifier is smaller than the pivot
		// we can extend the teritory of smallers by increment `smaller` by 1
		// and swap the qualifier and smaller, so we can get the new protential pivot position
		if nums[qualifier] < nums[pivot] {
			// increment the smaller by 1 and swap the qualifier to the new protential pivot position
			smaller++
			swap(nums, qualifier, smaller)
		}
	}
	// set the pivot to the right position - to the end of the smaller teritory (the right most)
	swap(nums, smaller, pivot)

	// update new pivot position
	pivot = smaller

	// sort the left and right part of the pivot
	quickSort(nums[:pivot])
	quickSort(nums[pivot+1:])
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
