package main

import (
	"fmt"
)

func main() {
	var input float32
	fmt.Println("Please print a float number:")
	fmt.Scan(&input)
	var trunc = int(input)
	fmt.Printf("Truncated number: %d \n", trunc)
}
