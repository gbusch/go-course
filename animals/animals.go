package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	for {
		fmt.Println("Please enter your request. Two strings in a single line, separated by space, first animal, then information.")
		fmt.Print(">")
		request := make([]string, 2)
		scanr := make([]interface{}, len(request))
		for i := range request {
			scanr[i] = &request[i]
		}
		fmt.Scanln(scanr...)
		var animal Animal
		if request[0] == "cow" {
			animal = Animal{food: "grass", locomotion: "walk", noise: "moo"}
		} else if request[0] == "bird" {
			animal = Animal{food: "worms", locomotion: "fly", noise: "peep"}
		} else if request[0] == "snake" {
			animal = Animal{food: "mice", locomotion: "slither", noise: "hsss"}
		} else {
			fmt.Println("wrong animal")
			continue
		}

		if request[1] == "eat" {
			animal.Eat()
		} else if request[1] == "move" {
			animal.Move()
		} else if request[1] == "speak" {
			animal.Speak()
		} else {
			fmt.Println("wrong request")
			continue
		}
	}
}
