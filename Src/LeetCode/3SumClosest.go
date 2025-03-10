package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(three_sum3_closest())
}
func three_sum3_closest() int {
	nums := []int{-1, 2, 1, -4}
	// nums := []int{0, 0, 0}
	target := 1
	result, min := 0, 0
	start, end := 0, len(nums)-1
	n := 2
	sort.Ints(nums)
	fmt.Println(nums)

	for start <= end {
		val := nums[start]
		i, j := start+1, end
		fmt.Println("------------------Start----------------------")
		for i < j && n > 0 {
			fmt.Printf("Pairs %d %d %d sum : %d, target : %d\n", val, nums[i], nums[j], result+val+nums[i]+nums[j], target)
			result = result + val + nums[i] + nums[j]

			if result > 0 {
				result = target - result
			} else if result < 0 {
				result = target - result
			}

			if min == 0 {
				min = result
			} else {
				if min > result && result >= 0 {
					min = result
				}
			}
			result = 0
			i++
			n--
		}
		fmt.Printf("Max Val : %d \n", min)
		start++
		n = 2
	}

	return min
}
