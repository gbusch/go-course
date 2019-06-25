package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)

	for {
		var input string
		fmt.Println("Please enter a number or 'X' to quit: ")
		fmt.Scan(&input)
		if input == "X" {
			break
		}
		i, err := strconv.Atoi(input)
		if err == nil {
			slice = append(slice, i)
			sort.Ints(slice)
			fmt.Println(slice)
		}
	}
	fmt.Println("Program ended with 'X'.")
}
