package sorting

func partition(arr []int, l int, h int) int {
	pivot := arr[l]
	i := l
	j := h

	for i < j {
		for i <= h && arr[i] <= pivot {
			i++
		}

		for j >= l && arr[j] > pivot {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[l], arr[j] = arr[j], arr[l]

	return j
}

func QuickSort(arr []int, l int, h int) {
	if l < h {
		pivot := partition(arr, l, h)

		QuickSort(arr, l, pivot-1)
		QuickSort(arr, pivot+1, h)
	}

}
