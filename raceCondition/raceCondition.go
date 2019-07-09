package main

import (
	"fmt"
)

func main() {
	var x int
	/* race condition = outcome depends on non-deterministic ordering (interleavings)
	Here we initialize an integer x (so it is 0). At the end of the function we increase the value,
	so it will be 1, and print it.
	In the middle, we add another function that should print the value.
	Depending on the ordering of the execution, this is either exectured before the increase of the number
	or afterwards. This result in two different results.*/
	go func() {
		fmt.Println(x)
	}()
	x++
	fmt.Println(x)
}
