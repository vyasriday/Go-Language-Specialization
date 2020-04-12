package main

import "fmt"

func main() {

	var unSortedArray []int
	var value int
	fmt.Println("Enter values one at a time")
	for i := 0; i < 10; i++ {
		fmt.Scan(&value)
		unSortedArray = append(unSortedArray, value)
	}

	bubbleSort(&unSortedArray)
	fmt.Println(unSortedArray)

}

func bubbleSort(array *[]int) {
	for i := 0; i < len((*array)); i++ {
		for j := 0; j < len((*array))-i-1; j++ {
			if (*array)[j] > (*array)[j+1] {

				swap(array, j)
			}
		}
	}

}

func swap(arrayPointer *[]int, index int) {
	temp := (*arrayPointer)[index+1]
	(*arrayPointer)[index+1] = (*arrayPointer)[index]
	(*arrayPointer)[index] = temp

}
