package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(t3_sum())
}
func t3_sum() [][]int {
	nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{-1, 0, 1, 0}
	sort.Ints(nums)

	result := [][]int{}
	seen := make(map[string]bool)

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				sub_array := []int{nums[i], nums[left], nums[right]}
				key := strconv.Itoa(nums[i]) + "," + strconv.Itoa(nums[left]) + "," + strconv.Itoa(nums[right])
				if !seen[key] {
					result = append(result, sub_array)
					seen[key] = true
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
