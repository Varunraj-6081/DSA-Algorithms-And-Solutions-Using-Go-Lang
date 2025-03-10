package main

import "fmt"

func main() {
	fmt.Println(twoSumArray())
}
func twoSumArray() []int {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := []int{}
	left, right := 0, len(nums)-1

	for left <= right {
		sum := nums[left] + nums[right]
		if sum == target {
			result = append(result, left+1)
			result = append(result, right+1)
			return result
		} else if sum <= target {
			left++
		} else {
			right--
		}
		fmt.Printf("Left Val :%d Rigth Val :%d \n", left, right)
	}
	return result
}
