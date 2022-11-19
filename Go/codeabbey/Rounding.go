package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	var results []float64
	fmt.Println("input data :")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		inp := scanner.Text()
		temp := strings.Fields(inp)
		a_ := temp[0]
		b_ := temp[1]
		a, _ := strconv.Atoi(a_)
		b, _ := strconv.Atoi(b_)

		a_float := float64(a)
		b_float := float64(b)

		result_ := a_float / b_float

		results = append(results, math.Round(result_))
	}

	s := forEach(results, fmt.Sprint)

	fmt.Println(strings.Join(s, " "))

}

func forEach(nums []float64, f func(...any) string) []string {
	var result []string
	for i := range nums {
		r := f(nums[i])

		result = append(result, r)
	}
	return result
}
