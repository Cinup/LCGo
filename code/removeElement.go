package code

func removeElement(nums []int, val int) int {
	i := 0
	n := len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]

			n--
		} else {
			i++
		}
	}
	return i
}
