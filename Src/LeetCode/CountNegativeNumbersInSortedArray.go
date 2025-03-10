package main

import "fmt"

func main() {
	fmt.Println(countNegativeNumbers())
}

func countNegativeNumbers() int {
	grid := [][]int{
		{3, 2},
		{-3, -3},
		{-3, -3},
		{-3, -3},
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("Nums[i,j] :%d, %d Value : %d  \n", i, j, grid[i][j])
			if grid[i][j] < 0 {
				count++
			}
		}
	}
	return count
}
