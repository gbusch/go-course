package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Please enter a string: ")
	fmt.Scanln(&input)
	lowered := strings.ToLower(input)
	fmt.Println(lowered)
	if strings.HasPrefix(lowered, "i") && strings.Contains(lowered, "a") && strings.HasSuffix(lowered, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
