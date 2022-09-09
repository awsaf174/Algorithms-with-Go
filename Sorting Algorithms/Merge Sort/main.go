package main

import "fmt"

func main() {
	arr := []int{2, 1, 3, 0}

	fmt.Println(MergeSort(arr))
}

func MergeSort(arr []int) (sortedArray []int) {

	var length int = len(arr)

	if length == 1 {
		return arr
	}

	splitIndex := length / 2

	leftArray := arr[:splitIndex]

	rightArray := arr[splitIndex:]

	return Merge(MergeSort(leftArray), MergeSort(rightArray))
}

func Merge(left []int, right []int) (mergedArray []int) {
	if len(left) == 1 && len(right) == 1 {
		if left[0] < right[0] {
			mergedArray = append(mergedArray, []int{left[0], right[0]}...)
		} else {
			mergedArray = append(mergedArray, []int{right[0], left[0]}...)
		}

		return mergedArray
	}
	j := 0
	i := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			mergedArray = append(mergedArray, left[i])
			i++
		} else {
			mergedArray = append(mergedArray, right[j])
			j++
		}
	}

	mergedArray = append(mergedArray, append(left[i:], right[j:]...)...)

	return mergedArray
}
