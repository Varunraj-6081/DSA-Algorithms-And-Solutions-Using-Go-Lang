package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(four4Sum())
}
func four4Sum() [][]int {
	// nums := []int{1, 0, -1, 0, -2, 2}
	// nums := []int{2, 2, 2, 2, 2}
	// nums := []int{0, 0, 0, 0, 0}
	nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{-3, -1, 0, 2, 4, 5}
	target := -1
	seen := make(map[string]bool)
	result := [][]int{}
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		fmt.Println("----------Start----------")
		for j := i + 1; j < len(nums)-1; j++ {
			k, l := j+1, len(nums)-1
			for k < l {
				fmt.Printf("i,j,k,l : %d,%d,%d,%d Sum : %d \n", nums[i], nums[j], nums[k], nums[l], nums[i]+nums[j]+nums[k]+nums[l])
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					key := strconv.Itoa(nums[i]) + "," + strconv.Itoa(nums[j]) + "," + strconv.Itoa(nums[k]) + "," + strconv.Itoa(nums[l])
					if !seen[key] {
						result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
						seen[key] = true
					}
					fmt.Println("----------******----------")
					fmt.Printf("Result : %d  %d,%d,%d %s\n", nums[i], nums[j], nums[k], nums[l], key)
					fmt.Println("----------******----------")
					k++
					l--
				} else if sum > target {
					l--
				} else {
					k++
				}
			}
		}
	}
	return result

}
