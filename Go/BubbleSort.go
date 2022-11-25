package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter the amount of numbers you want to sort: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// get the inputs from the user
	input := scanner.Text()

	// Convert the input to a slice of interger

	// split the input by spaces
	inputList := strings.Split(input, " ")

	// convert the string slice to an integer slice
	inputListInt := make([]int, len(inputList))

	for i, v := range inputList {
		inputListInt[i], _ = strconv.Atoi(v)
	}

	// call the bubbleSort function
	passes, swaps := bubbleSort(inputListInt)

	// print the number of passes and swaps
	fmt.Println(passes, swaps)

}

func generateRandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func bubbleSortTest(arr []int) {
	bubbleSort(arr)

	// join the slice with spaces
	arrString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), " "), "[]")

	fmt.Println(arrString)

}

func bubbleSort(arr []int) (int, int) {

	passes := 0
	swaps := 0

	isSorted := false

	if len(arr) == 0 {
		return passes, swaps
	}

	for !isSorted {
		isSorted = true
		passes += 1
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
				swaps += 1
				isSorted = false
			}
		}
	}

	return passes, swaps

}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
