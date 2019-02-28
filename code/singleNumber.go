package code

func singleNumber(nums []int) int {
	result := 0
	len := len(nums)
	for i := 0; i < len; i++ {
		result^=nums[i]
	}
	return  result
}
