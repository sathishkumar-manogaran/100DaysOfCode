package main

import "fmt"

func maxSubArray(nums []int) int {

	maxSum, currSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = Max(currSum+nums[i], nums[i])
		maxSum = Max(maxSum, currSum)
	}

	return maxSum
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	output := maxSubArray(nums)
	fmt.Println(output)
}
