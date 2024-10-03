package main

import (
	"fmt"
	"slices"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	order := make(map[int]int)
	for i, v := range arr2 {
		order[v] = i
	}

	var in_order []int
	var out_order []int

	for _, v := range arr1 {
		if slices.Contains(arr2, v) {
			in_order = append(in_order, v)
		} else {
			out_order = append(out_order, v)
		}
	}

	slices

}

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}
