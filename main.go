package main

import (
	"fmt"
	sortalgorithm "go-test/sort-algorithm"
)

func main() {
	arr := []int{4, 6, 2, 1, 7, 9, 5, 8, 3}
	arr = sortalgorithm.InsertionSort(arr)
	fmt.Println(arr)
}
