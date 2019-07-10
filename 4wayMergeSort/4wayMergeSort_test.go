package main

import "testing"

func TestMerge(t *testing.T) {
	tables := []struct {
		ldata  []int
		rdata  []int
		merged []int
	}{
		{[]int{1}, []int{}, []int{1}},
		{[]int{}, []int{1}, []int{1}},
		{[]int{1, 3}, []int{2, 4}, []int{1, 2, 3, 4}},
		{[]int{4, 8, 9}, []int{1, 5, 7}, []int{1, 4, 5, 7, 8, 9}},
	}

	for _, table := range tables {
		merged := Merge(table.ldata, table.rdata)
		for i := 0; i < len(merged); i++ {
			if merged[i] != table.merged[i] {
				t.Errorf("Merge failed. Element %d. Got: %d, expected: %d.", i, merged[i], table.merged[i])
			}
		}
	}
}

func TestSort(t *testing.T) {
	tables := []struct {
		data   []int
		sorted []int
	}{
		{[]int{9, 5, 7, 6, 1, 8, 2, 3}, []int{1, 2, 3, 5, 6, 7, 8, 9}},
	}

	for _, table := range tables {
		channel := make(chan []int)
		go Sort(table.data, channel)
		sorted := <-channel
		for i := 0; i < len(sorted); i++ {
			if sorted[i] != table.sorted[i] {
				t.Errorf("Sorting failed. Element %d. Got: %d, expected: %d.", i, sorted[i], table.sorted[i])
			}
		}
	}
}

func TestMergeSort(t *testing.T) {
	tables := []struct {
		data   []int
		sorted []int
	}{
		{[]int{9, 5, 7, 6, 1, 8, 2, 3}, []int{1, 2, 3, 5, 6, 7, 8, 9}},
	}

	for _, table := range tables {
		channel := make(chan []int)
		go MergeSort(table.data, channel)
		sorted := <-channel
		for i := 0; i < len(sorted); i++ {
			if sorted[i] != table.sorted[i] {
				t.Errorf("Sorting failed. Element %d. Got: %d, expected: %d.", i, sorted[i], table.sorted[i])
			}
		}
	}
}
