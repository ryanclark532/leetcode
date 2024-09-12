package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
  i, j := 0, len(height) -1
  water := 0

  for i < j {
    x := (float64(j-i) * math.Min(float64(height[i]), float64(height[j])))
    water = int(math.Max(float64(water),x ))
    if height[i] < height[j] {
      i +=1
    } else {
      j -=1
    }
  }


  return water
}



func main(){
fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
