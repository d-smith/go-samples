package sorting

func InsertionSort(a []int) {
	for i := 0; i < len(a); i++ {
		x := a[i]
		j := i - 1
		for j >= 0 && a[j] > x {
			a[j+1] = a[j]
			j = j - 1
		}
		a[j+1] = x
	}
}
