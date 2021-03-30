package main

func removeElement(nums []int, val int) int {
	if len(nums) == 1 {
		return 1
	}
	//length, i := 0, 0
	////i := 0
	//for i < len(nums) {
	//	if nums[i] == val {
	//		continue
	//	}
	//	nums[length] = nums[i]
	//	length++
	//}

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...) // append is in-built function
			i--
		}
	}

	return len(nums)
}

func main() {
	nums := []int{3, 2, 2, 3}
	removeElement(nums, 3)
}
