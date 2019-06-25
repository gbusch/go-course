package main

import (
	"fmt"
	"strings"
)

func FindIan(s string) bool {
	lowered := strings.ToLower(s)
	return strings.HasPrefix(lowered, "i") && strings.Contains(lowered, "a") && strings.HasSuffix(lowered, "n")
}

func main() {
	var input string
	fmt.Println("Please enter a string: ")
	fmt.Scanln(&input)
	if FindIan(input) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
