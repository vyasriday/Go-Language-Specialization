package main

import "fmt"

func main() {
	var input float32
	fmt.Println("Enter the floating point number")
	fmt.Scan(&input)
	truncated := int32(input)
	fmt.Println(truncated)
}
