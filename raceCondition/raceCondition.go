package main

import (
	"fmt"
)

func getNumber() int {
	var i int
	/* race condition = outcome depends on non-deterministic ordering (interleavings)
	here, we initialized i, so it is 0. At the end of the function we return it.
	In the middle of the function, we call another function that sets i to 42.
	Depending on the ordering of the execution, the value of i is either set to 42 and then returned
	(so it will print 42 in the main function)
	or it is returned already before the value is changed to 42
	(in this case it will print 0 in the main function)*/
	go func() {
		i = 42
	}()
	return i
}

func main() {
	fmt.Println(getNumber())
}
