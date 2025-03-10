package main

import "fmt"

func main() {
	fmt.Println(removeElement())
}

func removeElement() int {
	nums := []int{3, 2, 2, 3}
	val := 3
	k := 0
	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}
	return k
}
