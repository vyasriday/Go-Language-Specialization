package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.ToLower(str)
	fmt.Print(strings.Index(str, "i"))
	if strings.Index(str, "i") == 0 && strings.Index(str, "n") == len(str)-1 && strings.Index(str, "a") != -1 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
