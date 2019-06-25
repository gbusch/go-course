package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	names := make([]Name, 0, 3)
	var fname string
	fmt.Printf("Please enter the file name:")
	fmt.Scan(&fname)
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		splitname := strings.Split(string(line), " ")

		name1 := Name{fname: splitname[0], lname: splitname[1]}
		names = append(names, name1)
	}

	for _, name := range names {
		fmt.Printf("first name: %s, last name: %s \n", name.fname, name.lname)
	}

}
