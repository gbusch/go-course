package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Swap swaps the i and i+1 element of a slice sli
func Swap(sli []int, i int) {
	if i >= 0 && i < len(sli)-1 {
		tmp := sli[i+1]
		sli[i+1] = sli[i]
		sli[i] = tmp
	}
}

// BubbleSort performs the simple Bubblesort algorithm
func BubbleSort(sli []int) {
	for n := len(sli); n > 1; n-- {
		for i := 0; i < n-1; i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
			}
		}
	}
}

// ConvertInput takes comma-separated input string and converts it to a slice of integers
func ConvertInput(input string) []int {
	split := strings.Split(input, ",")
	slice := make([]int, 0, 1)
	for num, i := range split {
		j, err := strconv.Atoi(strings.TrimSpace(i))
		if err != nil {
			panic(err)
		}
		if num > 9 {
			fmt.Println("Use only up to 10 integers. Break here...")
			break
		}
		slice = append(slice, j)
	}
	return slice
}

func main() {
	var input string
	fmt.Println("Please enter up to 10 integers, separated by comma: ")
	fmt.Scanln(&input)
	s := ConvertInput(input)
	fmt.Println("Got the following slice of integers: ")
	fmt.Println(s)
	fmt.Println("Performing simple bubblesort algorithm...")
	BubbleSort(s)
	fmt.Println("Sorted slice: ")
	fmt.Println(s)
}
