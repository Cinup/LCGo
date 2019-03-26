package code

func twoSum1(nums []int, target int) []int {
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

func twoSum2(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		result := numbers[i] + numbers[j]
		switch {
		case result == target:
			return []int{i + 1, j + 1}
		case result < target:
			i++
		case result > target:
			j--
		}
	}
	return nil
}
