package main

import "testing"

func TestSwap(t *testing.T) {
	slice := []int{1, 2, 3}
	swapped := []int{1, 3, 2}
	Swap(slice, 1)
	for i := 0; i < len(slice); i++ {
		if slice[i] != swapped[i] {
			t.Errorf("Swap failed. Element %d. Got: %d, expected: %d.", i, slice[i], swapped[i])
		}
	}
}

func TestConvertInput(t *testing.T) {
	input := "3,2,1, 5"
	slice := ConvertInput(input)
	sliced := []int{3, 2, 1, 5}
	for i := 0; i < len(slice); i++ {
		if slice[i] != sliced[i] {
			t.Errorf("Conversion failed. Element %d. Got: %d, expected: %d.", i, slice[i], sliced[i])
		}
	}
}

func TestBubbleSort(t *testing.T) {
	slice := []int{4, 1, 7, 3, -1}
	sorted := []int{-1, 1, 3, 4, 7}
	BubbleSort(slice)
	for i := 0; i < len(slice); i++ {
		if slice[i] != sorted[i] {
			t.Errorf("BubbleSort failed. Element %d. Got: %d, expected: %d.", i, slice[i], sorted[i])
		}
	}
}
