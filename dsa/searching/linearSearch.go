package searching

func LinearSearch(arr []int, elem int) (int, bool) {

	var found bool = false

	var index int = -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			index = i
			found = true
			break
		}
	}

	return index, found
}
