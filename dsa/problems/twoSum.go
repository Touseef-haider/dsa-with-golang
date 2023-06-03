package problems

import "fmt"

func TwoSum() {
	fmt.Println("Two sum")

	var arr []int = []int{3, 4, 2, 1}

	var indexes []int = []int{}

	var target int = 5

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				indexes = append(indexes, i, j)
				break
			}
		}
	}

	// var i int
	// var j int = 1

	// for i < len(arr) {

	// 	if j >= len(arr)-1 && i <= len(arr)-1 {
	// 		i++
	// 		j = i + 1
	// 	}

	// 	if i == len(arr)-1 {
	// 		break

	// 	}

	// 	if arr[i]+arr[j] == target {
	// 		indexes = append(indexes, i, j)
	// 	}
	// 	j++
	// }

	fmt.Println(indexes)

}

// The time complexity for this algorithm is O(n^2)

// Is it possible to do this in O(n)?
// yes the second work use two pointers i and j, basically there is concept of sliding window which is being used here
