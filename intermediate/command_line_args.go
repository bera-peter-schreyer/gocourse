package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// os.Args is a slice of strings that contains the command-line arguments
	// os.Args[0] is the name of the program
	// os.Args[1:] are the command-line arguments passed to the program

	// Print the command-line arguments
	for i, arg := range os.Args {
		fmt.Printf("Argument %d: %s\n", i, arg)
	}

	// Define flags
	var name string
	var age int
	var isMale bool

	flag.StringVar(&name, "name", "John Doe", "Your name")
	flag.IntVar(&age, "age", 30, "Your age")
	flag.BoolVar(&isMale, "isMale", false, "Your gender")
	flag.Parse()

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Is Male: %t\n", isMale)
}	