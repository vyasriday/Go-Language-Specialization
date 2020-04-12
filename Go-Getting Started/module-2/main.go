// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var apple int
	fmt.Println("How many apples are there?")
	num, err := fmt.Scan(&apple);
	fmt.Println(num , err)
	fmt.Println(apple)
 }

