package utils

func Swap(arr []int, i int, j int) {

	var temp = arr[i]
	arr[i] = arr[j]
	arr[j] = temp

}
