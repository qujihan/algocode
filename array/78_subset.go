package main

func subsets(nums []int) [][]int {
	res := [][]int{}
	len := len(nums)
	for i := 0; i < (1 << len); i++ {
		temp := []int{}
		for j := 0; j < len; j++ {
			if (i & (1 << j)) != 0 {
				temp = append(temp, nums[j])
			}
		}
		res = append(res, temp)
	}

	return res
}
