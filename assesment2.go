package main

import "fmt"

func numRollsToTarget(n int, k int, target int) int {

  total :=0
  var nums []int
  for i:=1; k+1 > i; i++{
    nums = append(nums, i) 
  }


  for m := 0;m<n; m++{
    i:=0
    x:=len(nums)-1

  for ; i != x ; {
    y := nums[i] + nums[x]
      fmt.Println(y)
    if(y == target){
      total++
      i++
    } else if y > target{
      x--
    } else {
      i++
    }
  }

  }

  return total
}

func main() {
	fmt.Println(numRollsToTarget(30,30,500))
}
