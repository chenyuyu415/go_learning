package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Println("Hello Go!", os.Args[1])
		os.Exit(0)
	} else {
		fmt.Println("More argues!", os.Args[1:])
		os.Exit(1)
	}
}
