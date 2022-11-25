package main

import "math"

func main() {

}

/**
 * Find the product of all the elements in the array except the element at the current index.
 * @param arr The array of integers
 * @return The array of products
 * @example product([1, 2, 3, 4, 5]) = [120, 60, 40, 30, 24]
 * @example product([3, 2, 1]) = [2, 3, 6]
 * @example product([-1, -2, -3, 4, -5]) = [120, 60, 40, 30, 24]
 * @example product([1, 1, 1, 1, 1]) = [1, 1, 1, 1, 1]
 * @example product([]) = []
 */
func partialProduct(arr []int) []int {

	// If the array is empty, return an empty array
	if len(arr) == 0 {
		return []int{}
	}

	// Count amount of negative numbers
	negativeCount := 0
	for _, num := range arr {
		if num < 0 {
			negativeCount++
		}
	}

	productOfAll := 1

	// zeroCount is the amount of zeros in the array
	zeroCount := 0

	// Multiply all the numbers in the array, except zero
	for i := 0; i < len(arr); i++ {
		// if the number is zero, skip it
		if arr[i] == 0 {
			// Increment the zero count
			zeroCount++
			continue
		}
		productOfAll *= arr[i]
	}

	// containsZero is a boolean that indicates whether the array contains a zero
	containsZero := false
	if zeroCount > 0 {
		containsZero = true
	}

	result := make([]int, len(arr))

	// If the array contains more than one zero, the result will be an array of zeros
	if zeroCount > 1 {
		return result
	}

	// Divide the product by each number in the array, and store the result in the result array
	for i := 0; i < len(arr); i++ {

		// If the number is zero, then the product is the product of all the numbers except itself
		if arr[i] == 0 {
			result[i] = productOfAll
			// If there is odd number of negative numbers, then the product will be negative

			if negativeCount%2 != 0 {
				result[i] = -1 * int(math.Abs(float64(result[i])))
			}

			continue
		}

		// If the arr contains zero, then the product is zero
		if containsZero {
			result[i] = 0
			continue
		}

		result[i] = productOfAll / arr[i]

		// If there is an odd amount of negative numbers, and current number is negative, make the result positive
		if negativeCount%2 != 0 && arr[i] < 0 {
			result[i] = int(math.Abs(float64(result[i])))
		}

		// If there is an even amount of negative numbers, and current number is negative, make the result negative
		if negativeCount%2 == 0 && arr[i] < 0 {
			result[i] = -1 * int(math.Abs(float64(result[i])))

		}

	}

	return result
}
