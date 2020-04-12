package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findian(s string) (res bool) {
	res = false
	if strings.HasPrefix(strings.ToLower(s), "i") {
		if strings.HasSuffix(strings.ToLower(s), "n") {
			if strings.Contains(strings.ToLower(s), "a") {
				res = true
			}
		}
	}
	return
}

func main() {
	var userInput, rawInput string

	stringReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a word: ")
	rawInput, _ = stringReader.ReadString('\n')
	userInput = strings.TrimLeft(strings.TrimRight(rawInput, "\"\n"), "\"")
	if findian(userInput) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
