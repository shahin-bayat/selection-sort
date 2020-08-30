/*
 sort array from smallest to largest
*/
package main

import "fmt"

func main() {
	arr := []int{54, -17, 10, 128, -110, 500, -54}
	fmt.Println("Unsorted Array", arr, "Sorted:", selectionSort(arr))
}

func selectionSort(arr []int) []int {
	var sortedArr []int

	for len(arr) > 0 {
		minIdx := findMin(arr)
		sortedArr = append(sortedArr, arr[minIdx])
		arr = pop(arr, minIdx)
	}

	return sortedArr
}

func findMin(arr []int) int {
	min := arr[0]
	minIdx := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			minIdx = i
		}
	}
	return minIdx
}

// does NOT keep the order
func pop(arr []int, index int) []int {
	arr[index], arr[len(arr)-1] = arr[len(arr)-1], arr[index]
	return arr[:len(arr)-1]
}
