package algorithm

func insertion_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}
