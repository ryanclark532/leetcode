package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	values := make(map[int]int)

	for i, v := range nums {
		values[v] = i
	}

	i := 0

	for ; i < len(nums); i++ {
		complement := target - nums[i]
		val, ex := values[complement]
		if ex && val != i {
			return []int{val, i}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
