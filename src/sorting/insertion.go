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

func insertionSort(a[] int) {
  for i := 0; i < len(a); i++ {
    x := a[i]
    j := i - 1
    for j >= 0 && a[j] > x {
      a[j + 1] = a[j]
      j = j - 1
    }
    a[j+1] = x
  }
}

func main() {

  nums := [] int {23,5,4,1,100, -100, 8,7,2,6,3,-1}
  fmt.Println(isSorted(nums))
  insertionSort(nums)
  fmt.Println(isSorted(nums))
}
