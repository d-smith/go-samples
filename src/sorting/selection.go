package sorting

import "fmt"
import su "sorting/sortutil"

func swap(a[] int, i int, j int) {
	tmp := a[j]
	a[j] = a[i]
	a[i] = tmp
}

func SelectionSort(a[] int)  {
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
	fmt.Println(su.IsSorted(nums))
	SelectionSort(nums)
	fmt.Println(su.IsSorted(nums))
}
