package main

import "fmt"

func searchInsert(nums []int, target int) int {

	// case 1
	if target == 0 {
		return 0
	}

	// case 2 and case 3 covered
	for i, num := range nums {
		if num == target || target < num {
			return i
		}
	}

	// case 3
	for i, num := range nums {
		if target < num {
			return i
		}
	}

	// case 4
	if nums[len(nums)-1] < target {
		return len(nums) + 1
	}

	return -1
}

func searchInsert2(nums []int, target int) int {
	for key, val := range nums {
		if val >= target {
			return key
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 3, 5, 6}
	output := searchInsert(nums, 0)
	output = searchInsert2(nums, 0)
	fmt.Println(output)
}
