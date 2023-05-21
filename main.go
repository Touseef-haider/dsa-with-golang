package main

import (
	"fmt"

	"main.go/dsa/searching"
)

func main() {
	fmt.Println("DSA")

	// sorting.SelectionSort()

	// stack.StackDS()

	var arr []int = []int{2, 42, 2, 3, 5, 8}

	// var i int = 0
	// var j int = len(arr) - 1

	// fmt.Println(i, j)

	// queue.QueueDS()

	// sorting.QuickSort(arr, i, j)

	index, found := searching.LinearSearch(arr, 42)

	fmt.Print(index, found)

	// fmt.Println(arr)
}
