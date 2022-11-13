package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	inps := strings.Fields(GetInput("input data:"))
	var nums []int
	for i := range inps {
		n, _ := strconv.Atoi(inps[i])
		nums = append(nums, n)
	}
	fmt.Println(reduce(nums, sum))
}

func reduce(arr []int, method func(int, int) int) int {
	result := 0
	for _, v := range arr {
		result = method(result, v)
	}
	return result
}

func sum(a, b int) int {
	return a + b
}

func GetInput(prompt string) string {
	fmt.Print(prompt)
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func GetIntInput(prompt string) int {
	rt, err := strconv.ParseInt(GetInput(prompt), 10, 32)
	if err != nil {
		panic(err)
	}
	return int(rt)
}

func GetFloatInput(prompt string) float64 {
	rt, err := strconv.ParseFloat(GetInput(prompt), 10)
	if err != nil {
		panic(err)
	}
	return rt
}
