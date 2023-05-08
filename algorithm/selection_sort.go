package algorithm

func selection_sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp

	}
}
