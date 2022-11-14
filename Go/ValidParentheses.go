package main

import (
	"fmt"
)

func main() {
	testcase := "["
	fmt.Println(isValid(testcase))

}

func isValid(s string) bool {
	if len(s) == 0 {
		// Empty string
		return true
	}

	openers := "([{"
	closers := ")]}"

	var stack Stack
	for index := range s {
		c := string(s[index])

		if isIn(c, openers) {
			stack.push(c)
		} else if isIn(c, closers) {

			if stack.isEmpty() {
				// Closing on unopened parenthesis
				return false
			}

			p := stack.pop()
			if getCorrespondingCloser(p, openers, closers) != c {
				// Closing prenthesis type unmatched
				fmt.Println("HERE")
				return false
			}
		}
	}
	return stack.isEmpty()

}

type Stack []string

func getCorrespondingCloser(c string, openers string, closers string) string {
	for i := range openers {
		if c == string(openers[i]) {
			return string(closers[i])
		}
	}
	return ""
}

func (s *Stack) pop() string {
	if len(*s) == 0 {
		panic("cannot pop empty stack")
	}
	index := len(*s) - 1
	rt := (*s)[index]
	(*s).remove(index)
	return rt
}

func (s *Stack) remove(index int) []string {
	*s = append((*s)[:index], (*s)[index+1:]...)
	return *s
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(element string) {
	*s = append(*s, element)
}

func isIn(c string, ls string) bool {
	for index := range ls {
		if c == string(ls[index]) {
			return true
		}
	}
	return false
}
