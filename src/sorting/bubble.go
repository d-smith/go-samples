package sorting

func bubbleSort(a []int) {
	for itemCount := len(a) - 1; ; itemCount-- {
		hasChanged := false
		for i := 0; i < itemCount; i++ {
			if a[i] > a[i+1] {
				tmp := a[i+1]
				a[i+1] = a[i]
				a[i] = tmp
				hasChanged = true
			}
		}
		if hasChanged == false {
			break
		}
	}
}
