package main

import "fmt"

func shuffle(nums []int, n int) []int {

	output := make([]int, len(nums)) // creating new array using exiting array length
	c := 0
	for i := 0; i < n; i++ {
		output[c] = nums[i]
		c++
		output[c] = nums[n+i]
		c++
	}
	return output
}

func shuffle1(nums []int, n int) []int {
	x := nums[:n] // first half of the array value
	y := nums[n:] // second half of the array value
	fmt.Println(x)
	fmt.Println(y)

	//output := make([]int, len(nums))

	var output []int

	for i := 0; i < len(x); i++ {
		output = append(output, x[i])
		output = append(output, y[i])
	}
	return output
}

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	output := shuffle(nums, 3)
	output1 := shuffle1(nums, 3)
	fmt.Println(output)
	fmt.Println(output1)
}
