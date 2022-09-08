package main

import "fmt"

func main() {

	arr := []int{20, 10, 5, 15, 30, 50, 0, 0, 1, 200, 111}

	fmt.Println(sort(arr))
}

func sort(arr []int) (sortedArray []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	return arr
}
