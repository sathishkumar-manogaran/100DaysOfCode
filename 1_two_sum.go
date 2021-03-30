package main

import "fmt"

func twoSum(nums []int, target int) []int {

	result := make([]int, 2)                  // Array initialized with 2 index
	mapToStoreIndexValue := make(map[int]int) // Map created with key (int) and value (int) pair

	for i, num := range nums {
		if idx, existed := mapToStoreIndexValue[target-num]; existed { // check the key is already present or not
			result[0] = idx // store the index of map
			result[1] = i   // store the current iteration value
			return result
		}
		mapToStoreIndexValue[num] = i // store the current iteration value
	}
	return result
}

/*
1. map created and inserted with each value as key and iteration index as value
	map(2,0)
	map(11,1)
	map(7,2)
	map(8,3)
	map(15,4)
2. While doing iteration, execute the condition (target - current value) i.e (26 - 15 = 11)
	Now, 11 is already present in map with index value.
3. So get that map value and current index the return
*/

func main() {
	nums := []int{2, 11, 7, 8, 15}
	target := 26
	fmt.Println(twoSum(nums, target))
}
