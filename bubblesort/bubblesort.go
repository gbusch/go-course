package main

import (
	"fmt"
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

func main() {
	s := make([]int, 0, 10)
	fmt.Println("Please enter 10 integers, hit enter after each: ")
	for i := 0; i < 10; i++ {
		var input int
		fmt.Printf("Integer no %d: ", i+1)
		fmt.Scanln(&input)
		s = append(s, input)
	}
	fmt.Println("Got the following slice of integers: ")
	fmt.Println(s)
	fmt.Println("Performing simple bubblesort algorithm...")
	BubbleSort(s)
	fmt.Println("Sorted slice: ")
	fmt.Println(s)
}
