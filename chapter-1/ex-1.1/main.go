// Modify the echo program to also print os.Args[0], the name of the command that invoked it.

package main

import (
	"fmt"
	"os"
)

func printArray(a []string) {
	for _, arg := range a {
		fmt.Printf("%s \t", arg)
	}
}

func main() {
	var args = os.Args
	var s []string
	// fmt.Print(args)
	for _, arg := range args[0:] {
		s = append(s, arg)
	}
	fmt.Printf("Command: %s\n", s[0])
	fmt.Printf("Arguments: \t")
	printArray(s[1:])
	fmt.Println()
}
