package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var nums []int // Elements

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter n : ")
	scanner.Scan() // ignore

	fmt.Print("Enter elements : ")
	scanner.Scan()
	input := scanner.Text()
	inps := strings.Fields(input)
	nums = reduceString(inps, func(s string) int {
		r, _ := strconv.ParseInt(s, 10, 0)
		return int(r)
	})

	fmt.Println(reduceInt(nums, sum))
}

func sum(a, b int) int {
	return a + b
}

func reduceInt(arr []int, method func(int, int) int) int {
	result := 0
	for _, v := range arr {
		result = method(result, v)
	}
	return result
}

func reduceString(arr []string, method func(string) int) []int {
	var result []int
	for _, v := range arr {
		result = append(result, method(v))
	}
	return result
}
