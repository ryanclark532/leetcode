package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	solution := [][]int{}
	if len(nums) < 3 {
		return solution
	}

	sort.Ints(nums)
	seen := make(map[[3]int]bool)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicates for i
		}
		j, k := i+1, len(nums)-1

		for j < k {
			res := nums[i] + nums[j] + nums[k]

			if res == 0 {
				triplet := [3]int{nums[i], nums[j], nums[k]}
				if !seen[triplet] {
					solution = append(solution, []int{nums[i], nums[j], nums[k]})
					seen[triplet] = true
				}
				j++
				k--
				// Skip duplicates for j and k
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if res > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return solution
}

func removeDuplicates(arrays [][]int) [][]int {
	// Use a map to track unique arrays
	uniqueMap := make(map[string]struct{})
	var uniqueArrays [][]int

	for _, array := range arrays {
		// Convert the array to a string key for comparison
		key := fmt.Sprintf("%v", array)

		if _, found := uniqueMap[key]; !found {
			uniqueMap[key] = struct{}{}
			uniqueArrays = append(uniqueArrays, array)
		}
	}

	return uniqueArrays
}

func main() {
	nums := []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}

/*
func makeDistinct(nums []int) []int {
	visited := make(map[int]bool, len(nums))
	res := make([]int, len(nums))
	for _, v := range nums {
		if !visited[v] {
			res = append(res, v)
			visited[v] = true
		}
	}
	return res
}
*/
