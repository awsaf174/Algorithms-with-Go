package main

import "fmt"

func main() {
	arr := []int{20, 10, 5, 15, 30, 50, 0, 0, 1, 200, 111}
	fmt.Println(sort(arr))
}

func sort(arr []int) (sortedArray []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	return arr
}
