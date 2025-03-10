package main

import (
	"fmt"
	"strconv"
)

type stack []int

func (s *stack) push(val int) {
	*s = append(*s, val)
}
func (s *stack) pop() int {
	lastIndex := len(*s) - 1
	res := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return res
}

func main() {
	fmt.Println(evaluateReversePolishNotation())
}
func evaluateReversePolishNotation() int {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	var s stack
	for _, nums := range tokens {
		fmt.Println(s)
		if nums == "+" {
			a, b := s.pop(), s.pop()
			s.push(a + b)
		} else if nums == "-" {
			a, b := s.pop(), s.pop()
			s.push(b - a)
		} else if nums == "*" {
			a, b := s.pop(), s.pop()
			s.push(a * b)
		} else if nums == "/" {
			a, b := s.pop(), s.pop()
			s.push(b / a)
		} else {
			val1, _ := strconv.Atoi(nums)
			s.push(val1)
		}
	}
	return s[0]
}
