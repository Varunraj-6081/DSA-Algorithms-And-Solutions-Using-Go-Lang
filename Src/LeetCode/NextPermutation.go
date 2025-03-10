package main

func main() {
	nextPermutation()
}
func nextPermutation() {
	nums := []int{1, 2, 3}
	mid := 0
	k := 1
	swap := 0
	found := false
	for i := len(nums) - 2; i >= 0; i-- {
		unit := nums[len(nums)-k]
		// fmt.Printf("unit : %d,Nums[i]: %d \n", unit, nums[i])
		if unit > nums[i] {
			j := i + 1
			swap = j
			mid = i
			for j <= len(nums)-1 {
				// fmt.Printf("unit : %d,Nums[i]: %d \n", nums[i], nums[j])
				if nums[i] < nums[j] {
					swap = j
				}
				j++
			}
			nums[i], nums[swap] = nums[swap], nums[i]
			found = true
			break
		}
		k++
	}

	// Reverse the Value
	j := 0
	if found {
		j = mid + 1
	}

	for i := j; i <= len(nums)-1; i++ {
		temp := i
		min := nums[i]
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[j] < min {
				min = nums[j]
				temp = j
			}
		}
		nums[i], nums[temp] = min, nums[i]
	}
}
