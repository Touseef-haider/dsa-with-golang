package sorting

import "fmt"

func SelectionSort() {
	var arr []int = []int{8, 4, 1, 5, 9, 2}

	for i := 0; i < len(arr)-1; i++ {
		var min int = i

		for j := i + 1; j < len(arr); j++ {

			if arr[j] < arr[min] {
				min = j
			}

		}
		var temp = arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}

	fmt.Println(arr)
}
