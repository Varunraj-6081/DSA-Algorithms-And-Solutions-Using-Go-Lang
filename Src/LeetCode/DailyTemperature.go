package main

import "fmt"

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
	// 739 LeetCode Problem Topic : Stack
	fmt.Println(DailyTemperatures())
}
func DailyTemperatures() []int {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := make([]int, len(temperatures))
	var s stack
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(s) > 0 && temperatures[s[len(s)-1]] <= temperatures[i] {
			s = s[:len(s)-1]
		}
		if len(s) > 0 {
			result[i] = s[len(s)-1] - i
		} else {
			result[i] = 0
		}
		s = append(s, i)
	}

	return result
}
