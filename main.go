package main

import "fmt"

func main() {
	arr := []int{3, 5, 6, 8, 9, 4, 1, 5, 23, 66, 12, 90, 52}
	fmt.Println(len(arr))
	selection_sort(arr)
	// buble_sort(arr)
	fmt.Println(arr)
}

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

func buble_sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		{
			for j := 0; j < len(arr)-i-1; j++ {
				if arr[j] > arr[j+1] {
					temp := arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = temp
				}
			}

		}
	}
}
