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

func swap(a[] int, i int, j int) {
	tmp := a[j]
	a[j] = a[i]
	a[i] = tmp
}

func sort(a[] int)  {
	n := len(a)
	for i := 0; i < n - 1; i++ {
		minIndex := i;
		for j := i + 1; j < n; j++ {
			if(a[j] < a[minIndex]) {
				minIndex = j
			}
		}

		if(minIndex != i) {
			swap(a,i, minIndex)
		}
	}
}

func main() {

	nums := [] int {5,4,1,8,7,2,6,3}
	fmt.Println(isSorted(nums))
	sort(nums)
	fmt.Println(isSorted(nums))
}