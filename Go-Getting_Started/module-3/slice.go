package main

import (
	"fmt"
	"strconv"
)

func sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

func main() {

	var slice []int
	var input string
	
	for ;; {
		fmt.Println("Enter the interger for sorting and X to exit")
		fmt.Scan(&input)
		if input != "X" {
			i, _ := strconv.Atoi(input) 
			slice = append(slice, i)
			fmt.Println(sort(slice)) 
		}	else {
			break;
		}
	}

}
