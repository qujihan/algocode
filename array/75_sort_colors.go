package main

func sortColors(nums []int) {
	for index, point_0, point_1 := 0, 0, 0; index < len(nums); index++ {
		if nums[index] == 0 {
			nums[point_0], nums[index] = nums[index], nums[point_0]
			if point_0 < point_1 {
				nums[point_1], nums[index] = nums[index], nums[point_1]
			}
			point_0++
			point_1++
		} else if nums[index] == 1 {
			nums[point_1], nums[index] = nums[index], nums[point_1]
			point_1++
		}
	}
}

func main() {

}
