package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	data := []int{3, 1, 8, 2, 5}

	// fmt.Println(GetFurtherGreaterIndices(0, data))
	// fmt.Println(Handler(data))

	seed := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(seed)
	// fmt.Println(r1.Intn(100), time.Now().UnixNano())

	test := []int{}
	for i := 0; i < 1000; i++ {
		test = append(test, r1.Intn(100))
	}
	// sort.Ints(test)
	fmt.Println(test)

	data = test

	for _, value := range solve(data) {
		fmt.Print(fmt.Sprint(data[value]) + " ")
	}

}

func solve(data []int) []int {
	max := 0
	var lgp []int
	for i := range data {
		p, err := GetLongestPath(i, data)
		check(err)
		if len(p) > max {
			max = len(p)
			lgp = p
		}
	}

	return lgp
}

// func GetLongestPathLength(index int)

func GetLongestPath(index int, arr []int) ([]int, error) {
	if OutOfBound(index, arr) {
		return []int{}, errors.New("out of bound")
	}

	fgi, err := GetFurtherGreaterIndices(index, arr)
	max := 0
	var lgp []int
	check(err)
	for _, value := range fgi {
		p, err := GetLongestPath(value, arr)
		check(err)
		if len(p) > max {
			max = len(p)
			lgp = p
		}
	}
	rt := []int{index}
	rt = append(rt, lgp...)
	return rt, nil
}

func GetFurtherGreaterIndices(index int, arr []int) ([]int, error) {
	if OutOfBound(index, arr) {
		return []int{}, errors.New("out of bound")
	}

	var rt []int
	value := arr[index]
	for i := index; i < len(arr); i++ {
		if arr[i] > value {
			rt = append(rt, i)
		}
	}
	return rt, nil
}
func OutOfBound(index int, arr []int) bool {
	if len(arr) <= index {
		return true
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
