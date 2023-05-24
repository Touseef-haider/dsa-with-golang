package searching

import "fmt"

// remember binarySearch is used on sorted array
func BinarySearch(arr []int, searchValue int) {

	l := 0
	h := len(arr) - 1

	mid := (l + h) / 2

	fmt.Println(mid)

	for l < h {
		if arr[mid] == searchValue {
			fmt.Println("Found at ", mid)
			break
		} else {
			if searchValue < arr[mid] {
				h = mid
			} else {
				l = mid
			}
		}

	}

	fmt.Println("Not found")

}
