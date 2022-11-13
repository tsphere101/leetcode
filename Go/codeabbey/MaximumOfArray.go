package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("input data: ")
	scanner.Scan()
	nums := reduceString(strings.Fields(scanner.Text()), func(s string) int {
		r, _ := strconv.ParseInt(s, 10, 0)
		return int(r)
	})
	fmt.Println(reduceInt(nums, max), reduceInt(nums, min))
}

func reduceString(arr []string, f func(string) int) []int {
	var result []int
	for _, v := range arr {
		result = append(result, f(v))
	}
	return result
}

// For min
func reduceInt(arr []int, f func(int, int) int) int {
	var result int
	for _, value := range arr {
		result = f(result, value)
	}
	return result
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
