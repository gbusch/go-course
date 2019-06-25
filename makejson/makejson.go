package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name1 string
	var address1 string
	fmt.Println("Please enter a name: ")
	fmt.Scanln(&name1)
	fmt.Println("Please enter an address: ")
	fmt.Scanln(&address1)
	p1 := map[string]string{"name": name1, "address": address1}
	barr, err := json.Marshal(p1)
	fmt.Print("Print the map: ")
	fmt.Println(p1)
	if err == nil {
		fmt.Print("Print the json: ")
		fmt.Println(string(barr))
	}
}
