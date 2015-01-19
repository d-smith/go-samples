package main

import "fmt"

func isSorted(nums[] int) bool {
  for i := 1; i < len(nums); i++ {
    if(nums[i] < nums[i - 1]) {
      return false
    }
  }
  return true
}


func bubbleSort(a[] int) {
  for itemCount := len(a) -1; ; itemCount-- {
    hasChanged := false
    for i := 0; i < itemCount; i++ {
      if(a[i] > a[i + 1]) {
        tmp := a[i+1]
        a[i+1] = a[i]
        a[i] = tmp
        hasChanged = true
      }
    }
    if(hasChanged == false) {
        break
    }
  }
}

func main() {

  nums := [] int {23,5,4,1,100, -100, 8,7,2,6,3,-1}
  fmt.Println(isSorted(nums))
  bubbleSort(nums)
  fmt.Println(isSorted(nums))
}
