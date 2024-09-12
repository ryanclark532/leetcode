package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	combined := append(nums1, nums2...)
	sort.Ints(combined)

  length := len(combined)

  if length % 2 == 1{
    return float64(combined[length/2])
  }

  mid1, mid2 := combined[length/2-1], combined[length/2]
  return float64(mid1+mid2) /2.0 
}
