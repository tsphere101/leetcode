package main

import "fmt"

func main() {

	fmt.Println("Hello, world")

	var strs []string = []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Println(result)

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var common string = strs[0]
	for _, v := range strs {
		common = getCommonPrefix(common, v)
	}
	return common
}

func getCommonPrefix(a, b string) string {
	s := ""
	min := 0

	if len(a) < len(b) {
		min = len(a)
	} else {
		min = len(b)
	}

	for i := 0; i < min; i++ {
		if a[i] == b[i] {
			s += string(a[i])
		} else {
			break
		}
	}
	return s

}
