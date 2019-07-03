package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (a *Cow) Eat()   { fmt.Println("grass") }
func (a *Cow) Speak() { fmt.Println("moo") }
func (a *Cow) Move()  { fmt.Println("walk") }

type Bird struct{}

func (a *Bird) Eat()   { fmt.Println("worms") }
func (a *Bird) Speak() { fmt.Println("peep") }
func (a *Bird) Move()  { fmt.Println("fly") }

type Snake struct{}

func (a *Snake) Eat()   { fmt.Println("mice") }
func (a *Snake) Speak() { fmt.Println("hsss") }
func (a *Snake) Move()  { fmt.Println("slither") }

func main() {
	var command, animal, action string
	animals := make(map[string]Animal, 0)
	for {
		fmt.Print("> ")
		fmt.Scanln(&command, &animal, &action)
		switch command {
		case "newanimal":
			switch action {
			case "cow":
				animals[animal] = &Cow{}
			case "bird":
				animals[animal] = &Bird{}
			case "snake":
				animals[animal] = &Snake{}
			default:
				fmt.Println("Error, animal type does not exist.")
				continue
			}
		case "query":
			a, ok := animals[animal]
			if ok {
				switch action {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				default:
					fmt.Println("Error, action dows not exist.")
					continue
				}
			} else {
				fmt.Println("Error, animal did not exist.")
			}
		}
	}
}
