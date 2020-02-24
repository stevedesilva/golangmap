package main

import (
	"fmt"
	"os"

	"github.com/stevedesilva/golangmap/mymap"
)

func main() {
	args := os.Args[1:]
	data, err := mymap.New().Students(args)
	if err != nil {
		fmt.Printf("Please type a Hogwarts house name.\n")
		return
	}

	if data != nil {
		fmt.Println("data:", data)
	} else {
		fmt.Printf("Sorry. I don't know anything about %q \n", args[0])
	}

}
