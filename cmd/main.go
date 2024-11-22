package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()

	default:
		panic("Invalid command input")
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func run() {
	fmt.Printf("Sticky running %v\n", os.Args[2:])
}
