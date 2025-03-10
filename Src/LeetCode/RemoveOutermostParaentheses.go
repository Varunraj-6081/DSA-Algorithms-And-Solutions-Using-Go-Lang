package main

import "fmt"

func main() {
	// 1021 LeetCode Topic : Stack
	fmt.Println(removeOuterParenthesis())
}
func removeOuterParenthesis() string {
	s := "(()())(())"
	result := ""
	count := 0
	for _, str := range s {
		value := string(str)
		if value == ")" {
			count--
		}
		if count > 0 {
			result = result + value
		}
		if value == "(" {
			count++
		}
		fmt.Println(result)
	}

	return result

}
