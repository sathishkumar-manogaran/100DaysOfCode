package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	i := 0
	j := 1
	for j < len(nums) {
		/*if nums[i] == nums[j] {
			j++
		} else if nums[i+1] == nums[j] {
			i++
			j++
		} else {
			nums[i+2] = nums[j]
			i++
			j++
		}*/

		if nums[i] == nums[j] {
			j++
		} else {
			nums[i+1] = nums[j]
			i++
			j++
		}
	}

	return i + 1
	//return i + 2
}

/*
take first value from the array and check with next value
if it match then just go to next element
else change the search index to +1 and increase all the index to next one
Note: We are not modifying anything in array, we are changing / increasing only index value to identify new length
*/

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	output := removeDuplicates(nums)
	fmt.Print(output)

}
