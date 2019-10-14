package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])

	// Include name of invoking command - Ex 1.1
	fmt.Println(strings.Join(os.Args[0:], " "))

	// Print index and value of args, 1 per line - Ex 1.2
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
