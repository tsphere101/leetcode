package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func toCel(num int) int {
	result := math.Round(float64(num-32) * 5.0 / 9.0)
	return int(result)
}

func main() {
	fmt.Println("input data :")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	converted := []int{}
	inp := strings.Fields(scanner.Text())
	inp = inp[1:]
	var inps_num []int

	for i := range inp {
		elem, _ := strconv.Atoi(inp[i])
		inps_num = append(inps_num, elem)
	}

	for i := range inps_num {
		converted = append(converted, toCel(inps_num[i]))
	}

	fmt.Println(func() string {
		var result string
		var elems []string
		for i := range converted {
			elems = append(elems, fmt.Sprint(converted[i]))
		}
		result = strings.Join(elems, " ")
		return result
	}())

}
