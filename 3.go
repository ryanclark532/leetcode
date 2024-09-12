package main

import(
  "fmt"
)


func lengthOfLongestSubstring(s string) int {
  checked := make(map[string]bool)  
  for i, c := range s {
    if _, ex := checked[string(c)]; ex {
      return i 
    } else {
      checked[string(c)] = true 
    }

  }
  return 1
}



func main(){
  fmt.Println(lengthOfLongestSubstring("abcabcbb"))
  fmt.Println(lengthOfLongestSubstring("bbbbbb"))
  fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
