package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	var target = 9
	ret := twoSum(nums, target)
	fmt.Println(ret)

}
