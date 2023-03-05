package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, index := m-1, n-1, len(nums1)-1; index >= 0; {
		if p1 < 0 {
			nums1[index] = nums2[p2]
			p2--
			index--
			continue
		}
		if p2 < 0 {
			return
		}
		if nums1[p1] > nums2[p2] {
			nums1[index] = nums1[p1]
			p1--
			index--
		} else {
			nums1[index] = nums2[p2]
			p2--
			index--
		}
	}

}
