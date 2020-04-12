package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Name struct
type Name struct {
	fname string
	lname string
}

func main() {

	var fileName string
	var Names []Name
	fmt.Println("Enter the name of the file")
	fmt.Scan(&fileName)
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	fileData := string(content)
	fileArray := strings.Split(fileData, "\n")
	for i := 0; i < len(fileArray); i++ {
		nameArray := strings.Split(string(fileArray[i]), " ")
		var temp Name
		temp.fname = nameArray[0]
		temp.lname = nameArray[1]
		Names = append(Names, temp)
	}

	for _, value := range Names {
		fmt.Println("First Name" , value.fname)
		fmt.Println("Last Name", value.lname)
	}

}
