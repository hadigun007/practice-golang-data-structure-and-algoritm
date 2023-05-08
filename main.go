package main

import "fmt"

func main() {
	arr := []int{3, 5, 6, 8}
	selection_sort(arr, 4)
	fmt.Println(arr)
}
func selection_sort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		iMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[iMin] {
				iMin = j
			}
			temp := arr[i]
			arr[i] = arr[iMin]
			arr[iMin] = temp
		}
	}
}
