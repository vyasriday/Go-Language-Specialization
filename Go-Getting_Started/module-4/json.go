package main

import (
	"fmt"
	"encoding/json"
	)
func main() {
	
	var name string
	var addr string

	fmt.Println("Enter your name!")
	fmt.Scan(&name)
	fmt.Println("Enter your address")
	fmt.Scan(&addr)
	mapPerson := make(map[string]string)
	mapPerson["name"] = name
	mapPerson["address"] = addr
	fmt.Println(mapPerson)
	barray, _ := json.Marshal(mapPerson)

	fmt.Println(string(barray))
}