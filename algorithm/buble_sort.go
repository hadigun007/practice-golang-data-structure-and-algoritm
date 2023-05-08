package algorithm

func buble_sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		{
			for j := 0; j < len(arr)-i-1; i++ {
				if arr[j] > arr[j+1] {

					temp := arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = temp
				}
			}

		}
	}
}
