package code
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		i, ok := m[target-num]
		if ok {
			return []int{i, index}
		}
		m[num] = index
	}
	return nil
}