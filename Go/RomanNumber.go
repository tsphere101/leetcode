package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		inp, err := GetInput("Enter Roman Number : ")

		if err != nil {
			fmt.Println("Error")
		}
		fmt.Println("result >", romanToInt(inp))
	}
}

func GetInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func romanToInt(s string) int {
	var ref map[string]int = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	if len(s) == 1 {
		return ref[strings.ToUpper(string(s[0]))]
	}

	var result int = 0
	value := 0
	next := 0
	for i := 0; i < len(s)-1; i++ {
		value = ref[strings.ToUpper(string(s[i]))]
		next = ref[strings.ToUpper(string(s[i+1]))]
		if next > value {
			result -= value
		} else {
			result += value
		}
	}
	result += next

	return result
}
