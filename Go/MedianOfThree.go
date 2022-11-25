package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// n is number of inputs
	n := 0

	// scan the number of inputs
	scanner := bufio.NewScanner(os.Stdin)

	// scan the number of inputs
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	// create a slice of size n
	arr := make([]int, n)

	// scan the inputs, 3 int at a time and put them in the slice
	for i := 0; i < n; i++ {
		fmt.Print("Enter 3 integers: ")
		scanner.Scan()
		subarr := make([]int, 3)

		// split the input into 3 integers
		inps := strings.Fields(scanner.Text())

		// convert the inps to integers
		for i, inp := range inps {
			subarr[i], _ = strconv.Atoi(inp)
		}

		// put the median of the 3 integers in the slice
		arr[i] = median(subarr)
	}

	// print the result, join the slice with spaces
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), " "), "[]"))

}

// find the median of slice
func median(arr []int) int {
	// sort the arr
	sort.Ints(arr)

	var middle int

	// find the middle index

	// if the length of arr is odd
	if len(arr)%2 == 1 {
		middle = len(arr) / 2
	}

	// if the length of arr is even
	if len(arr)%2 == 0 {
		middle = len(arr) / 2
	}

	// return the value at the middle index
	return arr[middle]

}
