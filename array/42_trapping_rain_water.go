package main

import (
	"fmt"
)

func trap(height []int) int {
	ans := 0
	left, right, left_max, right_max := 0, len(height)-1, 0, 0
	for left < right {
		if height[right] > height[left] {
			if height[left] > left_max {
				left_max = height[left]
			} else {
				ans = ans + (left_max - height[left])
				left++
			}
		} else {
			if height[right] > right_max {
				right_max = height[right]
			} else {
				ans = ans + (right_max - height[right])
				right--
			}
		}
	}
	return ans
}

func main() {

	height := []int{4, 2, 0, 3, 2, 5}
	ans := trap(height)
	fmt.Println(ans)
}
