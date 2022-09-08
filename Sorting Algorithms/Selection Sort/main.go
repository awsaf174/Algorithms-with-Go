package main

import "fmt"

func main() {
	arr := []int{20, 10, 5, 15, 30, 50, 0, 0, 1, 200, 111}
	fmt.Println(sort(arr))
}

func sort(arr []int) (sortedArray []int) {

	//smallestNumber := 0

	for i := 0; i < len(arr); i++ {
		smallestNumber := arr[i]
		index := i
		for j := i + 1; j < len(arr); j++ {
			if smallestNumber > arr[j] {
				smallestNumber = arr[j]
				index = j
			}
		}
		arr[index] = arr[i]
		arr[i] = smallestNumber

	}

	return arr
}
