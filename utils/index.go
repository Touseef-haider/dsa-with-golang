package utils

func Swap(arr []int, i int, j int) {

	var temp = arr[i]
	arr[i] = arr[j]
	arr[j] = temp

}

func ElementExistsInArray[T comparable](arr []T, el T) bool {

	found := false

	for _, v := range arr {

		if v == el {
			found = true
			break
		}
	}

	return found

}

func RemoveIndex[T comparable](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)

}
