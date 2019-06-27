package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn returns a function that takes time and gives displacement.
// Paramters are acceleration, start velocity and start displacenemt
func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
	return fn
}

// ReadParameters is a helper function for reading the parameters from user input
func ReadParameters() []float64 {
	var param []float64
	fmt.Println("Please enter acceleration: ")
	fmt.Scanln(&param[0])
	fmt.Println("Please enter initial velocity: ")
	fmt.Scanln(&param[1])
	fmt.Println("Please enter initial displacement: ")
	fmt.Scanln(&param[2])
	return param
}

func main() {
	param := ReadParameters()
	fn := GenDisplaceFn(param[0], param[1], param[2])
	var time float64
	fmt.Println("Please enter time: ")
	fmt.Scanln(&time)
	fmt.Print("The displacement is: ")
	fmt.Println(fn(time))
}
