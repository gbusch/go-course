package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Merge(ldata []int, rdata []int) (result []int) {
	result = make([]int, len(ldata)+len(rdata))
	lidx := 0
	ridx := 0
	for lidx < len(ldata) && ridx < len(rdata) {
		if ldata[lidx] <= rdata[ridx] {
			result[lidx+ridx] = ldata[lidx]
			lidx++
		} else {
			result[lidx+ridx] = rdata[ridx]
			ridx++
		}
	}
	for lidx < len(ldata) {
		result[lidx+ridx] = ldata[lidx]
		lidx++
	}
	for ridx < len(rdata) {
		result[lidx+ridx] = rdata[ridx]
		ridx++
	}
	return
}

func Sort(data []int, r chan []int) {
	if len(data) == 1 {
		r <- data
		return
	}

	channel := make(chan []int, 2)
	middle := len(data) / 2

	go Sort(data[:middle], channel)
	go Sort(data[middle:], channel)

	ldata := <-channel
	rdata := <-channel
	close(channel)

	r <- Merge(ldata, rdata)
	return
}

func MergeSort(data []int, r chan []int) {
	if len(data) < 4 {
		panic("list must have length >= 4")
	}

	channel := make(chan []int, 4)
	quart := len(data) / 4
	part1 := data[:quart]
	part2 := data[quart : 2*quart]
	part3 := data[2*quart : 3*quart]
	part4 := data[3*quart:]

	fmt.Print("Subarray 1:")
	fmt.Println(part1)
	fmt.Print("Subarray 2:")
	fmt.Println(part2)
	fmt.Print("Subarray 3:")
	fmt.Println(part3)
	fmt.Print("Subarray 4:")
	fmt.Println(part4)

	go Sort(part1, channel)
	go Sort(part2, channel)
	go Sort(part3, channel)
	go Sort(part4, channel)

	data1 := <-channel
	data2 := <-channel
	data3 := <-channel
	data4 := <-channel

	close(channel)

	merge1 := Merge(data1, data2)
	merge2 := Merge(data3, data4)

	r <- Merge(merge1, merge2)
	return
}

// ConvertInput takes comma-separated input string and converts it to a slice of integers
func ConvertInput(input string) []int {
	split := strings.Split(input, ",")
	slice := make([]int, 0, 1)
	for _, i := range split {
		j, err := strconv.Atoi(strings.TrimSpace(i))
		if err != nil {
			panic(err)
		}
		slice = append(slice, j)
	}
	return slice
}

func main() {
	var input string
	fmt.Println("Please enter at least 4 integers, separated by comma, spaces are not allowed. ")
	fmt.Println("For example: 3,2,6,4,1,9,6,3")
	fmt.Println("Not allowed: 3, 2, 6, 4, 1, 9, 6, 3")
	fmt.Scanln(&input)
	s := ConvertInput(input)

	result := make(chan []int)
	go MergeSort(s, result)

	r := <-result
	fmt.Println("Sorted slice of integers:")
	for _, v := range r {
		fmt.Printf("%d, ", v)
	}
	close(result)
	fmt.Println()
}
