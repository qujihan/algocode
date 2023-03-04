package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for index, value := range nums {
		if value > 0 {
			break
		}

		if index > 0 && nums[index] == nums[index-1] {
			continue
		}

		left, right := index+1, len(nums)-1
		for left < right {
			sum := nums[index] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[index], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}

func main() {
	nums := []int{0, 1, 1}
	ans := threeSum(nums)
	fmt.Println(ans)

}
