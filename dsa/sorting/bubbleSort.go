package sorting

import "fmt"

func BubbleSort() {

	var arr []int = []int{4, 3, 7, 15}

	// i=0
	// 4 < 2 No
	// 13 < 4 No
	// 7 < 13 swap
	// [2,4,7,13,9]
	// 9 < 13 swap
	// [2,4,7,9,13]

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				var temp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}

	fmt.Println(arr)
}
