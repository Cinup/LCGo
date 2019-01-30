package code

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m] = nums2[i]
		m++
	}
	sort.Ints(nums1)
}
