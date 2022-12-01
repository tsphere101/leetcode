package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func fib(num int) *big.Int {
	memory := []*big.Int{new(big.Int).SetInt64(0), new(big.Int).SetInt64(1), new(big.Int).SetInt64(1)}
	var i int
	for i = 2; i < num; i++ {

		memory = append(memory, new(big.Int))
		x := memory[i+1]

		a := memory[i]
		b := memory[i-1]
		x.Add(a, b)
	}
	return memory[num]
}

func findFib(n *big.Int) int {
	i := 0
	for {
		fib_ := fib(i)
		if fib_.Cmp(n) == 0 {
			return i
		}
		i++
	}
}

func intSliceToString(arr []int) string {
	var stringSlice []string
	for _, value := range arr {
		stringSlice = append(stringSlice, fmt.Sprint(value))
	}
	return strings.Join(stringSlice, " ")
}

func main() {

	fmt.Println("input data:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	inps := []*big.Int{}
	for i := 0; i < n; i++ {
		scanner.Scan()
		// num , _ := strconv.Atoi(scanner.Text())
		num, _ := new(big.Int).SetString(scanner.Text(), 10)
		inps = append(inps, num)
	}

	result := []int{}
	for i := range inps {
		result = append(result, findFib(inps[i]))
	}

	fmt.Println(intSliceToString(result))

}
