package sorting

import "fmt"

func SelectionSort() {

	var arr []int = []int{4, 3, 7, 1, 5}

	for i := 1; i < len(arr); i++ {

		var temp int = arr[i]
		var j int = i - 1

		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--

		}

		arr[j+1] = temp

	}

	fmt.Println(arr)

}
