package sortutil

func IsSorted(nums[] int) bool {
  for i := 1; i < len(nums); i++ {
    if(nums[i] < nums[i - 1]) {
      return false
    }
  }
  return true
}
