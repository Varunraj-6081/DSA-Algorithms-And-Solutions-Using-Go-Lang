package main

import "fmt"

func main() {
	fmt.Println(container_most_water())
}
func container_most_water() int {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// height := []int{1, 1}

	left, right := 0, len(height)-1
	max_value, cur_val := 0, 0

	for left <= right {
		if height[left] <= height[right] {
			width := right - left
			cur_val = width * height[left]

			if max_value == 0 {
				max_value = cur_val
			} else {
				if cur_val > max_value {
					max_value = cur_val
				}
			}

			left = left + 1
		} else if height[left] > height[right] {
			width := right - left
			cur_val = width * height[right]

			if max_value == 0 {
				max_value = cur_val
			} else {
				if cur_val > max_value {
					max_value = cur_val
				}
			}
			right = right - 1
		}

	}

	return max_value
}
