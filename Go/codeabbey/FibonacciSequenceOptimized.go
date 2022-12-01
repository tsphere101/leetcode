package main

import (
	"fmt"
	"math/big"
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

func findFib2() int {

	var fibSeq map[string]int
	fibSeq = make(map[string]int)

	fibSeq["0"] = 0
	fibSeq["1"] = 1
	fibSeq["1"] = 2

	a, b := big.NewInt(1), big.NewInt(2)
	c := big.NewInt(0)
	for i := 2; i < 10; i++ {
		c.Add(a, b)
		fibSeq[c.String()] = i
		a, b = c, a
	}
	fmt.Println(fibSeq["34"])

	return 1
}

func intSliceToString(arr []int) string {
	var stringSlice []string
	for _, value := range arr {
		stringSlice = append(stringSlice, fmt.Sprint(value))
	}
	return strings.Join(stringSlice, " ")
}

func main() {
	findFib2()
}

// func main_() {

// 	fmt.Println("input data:")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	n, _ := strconv.Atoi(scanner.Text())
// 	inps := []*big.Int{}
// 	for i := 0; i < n; i++ {
// 		scanner.Scan()
// 		// num , _ := strconv.Atoi(scanner.Text())
// 		num, _ := new(big.Int).SetString(scanner.Text(), 10)
// 		inps = append(inps, num)
// 	}

// 	result := []int{}
// 	for i := range inps {
// 		result = append(result, findFib(inps[i]))
// 	}

// 	fmt.Println(intSliceToString(result))

// }
