package code

func searchInsert(nums []int, target int) int {
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
	}
	return i
}
