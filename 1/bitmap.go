package main

import (
	"fmt"
)

func BitMapSort(max int, arr []int) (result []int) {
	var bitmap [101]byte
	fmt.Println(bitmap)
	for _, v := range arr {
		bitmap[v] = 1
	}
	for index, i := range bitmap {
		if i == 1 {
			fmt.Println(i, index)
		}
	}
	return
}

func bubblingSort(arr []int) (result []int) {
	size := len(arr)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		//fmt.Println(arr)
	}
	return arr
}

func main() {
	r := BitMapSort(101, []int{10, 100, 99, 80, 20, 40, 56, 3})
	fmt.Println(r)
	fmt.Println("------------------------")
	fmt.Println(bubblingSort([]int{10, 100, 99, 80, 20, 40, 56, 3, 200}))
}
