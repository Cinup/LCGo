package code

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	result := 0
	for i := 0; i < len(nums); i = i + 2 {
		result += nums[i]
	}
	return result
}
