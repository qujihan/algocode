package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}

	if len == 1 {
		return nums[0]
	}

	opt := make([]int, len)
	opt[0] = nums[0]
	res := nums[0]
	for index := 1; index < len; index++ {
		if opt[index-1] < 0 {
			opt[index] = nums[index]
		} else {
			opt[index] = opt[index-1] + nums[index]
		}
		if opt[index] > res {
			res = opt[index]
		}
		// fmt.Println("at index:", index, " res is:", res)

	}

	return res
}

func main() {
	nums := []int{5, 4, -1, 7, 8}
	res := maxSubArray(nums)
	fmt.Println(res)
}
